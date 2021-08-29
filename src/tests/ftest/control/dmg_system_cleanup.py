#!/usr/bin/python
"""
  (C) Copyright 2020-2021 Intel Corporation.

  SPDX-License-Identifier: BSD-2-Clause-Patent
"""
from apricot import TestWithServers
from avocado.core.exceptions import TestFail
from command_utils import CommandFailure
from general_utils import DaosTestError
from pydaos.raw import (DaosApiError, DaosPool)


class DmgSystemCleanupTest(TestWithServers):
    """Test Class Description:
    This test covers the following requirement.
    (SRS-326) Pool Management - automatic pool handle revocation

    Test step:
    1. Create 2 pools.
    2. Create a container in each pool.
    3. Call dmg system cleanup host.
    4. Call daos pool list-cont on each pool. Each should fail with -1012.

    :avocado: recursive
    """

    def test_dmg_system_cleanup_one_host(self):
        """
        JIRA ID: DAOS-6471

        Test Description: Test dmg system cleanup.

        :avocado: tags=all,full_regression
        :avocado: tags=small,dmg
        :avocado: tags=control,dmg_system_cleanup
        """
        # Create 2 pools and create a container in each pool.
        self.pool = []
        self.container = []
        for _ in range(2):
            self.pool.append(self.get_pool())
            self.container.append(self.get_container(self.pool[-1]))

        # Create 5 more connections to each pool
        self.pool_handles = []
        for pool in self.pool:
            for _ in range(5):
                handle = self.get_pool(create=False, connect=False)
                handle.pool = DaosPool(self.context)
                handle.uuid = pool.uuid
                connected = handle.connect(2)
                self.pool_handles.append(handle)

        # Check to make sure we can access the pool
        try:
            for i in range(2):
                self.container[i].get_params(self)
                self.container[i].write_objects()
        except (DaosApiError, DaosTestError) as error:
                self.fail("Unable to write container #{}: {}".format(i, error))

        # Call dmg system cleanup on the host and create cleaned pool list.
        dmg_cmd = self.get_dmg_command()
        result = dmg_cmd.system_cleanup(self.hostlist_servers[0], verbose=True)

        # Build list of pools and how many handles were cleaned (should be 6 each)
        actual_handle_counts = dict()
        for pool in result["response"]["pools"]:
            actual_handle_counts[pool["pool_id"].lower()] = pool["count"]
        # Attempt to access the pool again (should fail)
        for i in range(2):
            try:
                self.container[i].write_objects()
            except (DaosApiError, DaosTestError, TestFail) as error:
                self.log.info("Unable to write container #{}: as expected {}".format(i, error))
            else:
                self.fail("Wrote to container #{} when it should have failed: {}".format(i, error))

        # Build a list of pool IDs and counts (6) to compare against
        # our cleanup results.
        expected_handle_count = dict()
        for pool in self.pool:
            expected_handle_count[pool.uuid.lower()] = 6

        # Clear pool and containter list to avoid trying to destroy them.
        self.pool = []
        self.container = []

        # Compare results
        self.assertEqual(len(expected_handle_count),len(actual_handle_counts), "Cleaned up handles does not match the expected amount.")
        for k in expected_handle_count:
            self.assertEqual(expected_handle_count[k], actual_handle_counts[k], "Count for {} is not equal: expected {}, actual {}".format(k,expected_handle_count[k],actual_handle_counts[k]))

        # Ensure that our set of expected and actual pools are the same
        self.log.info("Test passed!")
