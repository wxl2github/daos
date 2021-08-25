#!/usr/bin/python
"""
  (C) Copyright 2020-2021 Intel Corporation.

  SPDX-License-Identifier: BSD-2-Clause-Patent
"""
from apricot import TestWithServers
from command_utils import CommandFailure


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

        daos_cmd = self.get_daos_command()

        # Call daos pool list-cont on each pool. It should succeed.
        for i in len(self.pool):
            try:
                daos_cmd.container_list(pool=self.pool[i].uuid)
                self.log.info(
                    "daos pool list-cont on pool[%d] succeeded as expected", i)
            except CommandFailure:
                self.fail(
                    "daos pool list-cont on pool[{}] failed unexpectedly!".format(i))

        # Build a list of pool IDs to compare against our cleanup results.
        expected_pool_list = set()
        for pool in self.pool:
            expected_pool_list.add(pool.uuid)

        # Call dmg system cleanup on the host and create cleaned pool list.
        dmg_cmd = self.get_dmg_command()
        result = dmg_cmd.system_cleanup(self.hostlist_clients, verbose=True)

        actual_pool_list = set()
        for pool in result["response"]["pools"]:
            actual_pool_list.add(pool["Id"])

        # Ensure that our set of expected and actual pools are the same
        self.assertEqual(len(expected_pool_list ^ actual_pool_list), 0)


        # Call daos pool list-cont on the pools again. It should fail.
        for i in len(self.pool):
            try:
                daos_cmd.container_list(pool=self.pool[i].uuid)
            except CommandFailure:
                self.log.info(
                    "daos pool list-cont on pool[%d] failed as expected", i)
            else:
                self.fail(
                    "daos pool list-cont on pool[{}] succeeded unexpectedly!".format(i))
        self.log.info("Test passed!")
