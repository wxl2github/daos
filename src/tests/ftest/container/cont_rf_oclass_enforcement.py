#!/usr/bin/python
"""
  (C) Copyright 2019-2021 Intel Corporation.

  SPDX-License-Identifier: BSD-2-Clause-Patent
"""
from container_rf_test_base import RbldContRedundancyFactor


class RbldContRedundancyFactorEnforce(RbldContRedundancyFactor):
    # pylint: disable=too-many-ancestors
    """Test Container redundancy factor enforcement with oclass traffic
       and rebuild.

    :avocado: recursive
    """

    def __init__(self, *args, **kwargs):
        """Initialize a Rebuild Container RF with ObjClass Write object."""
        super().__init__(*args, **kwargs)
        self.daos_cmd = None

    def test_container_redundancy_factor_oclass_enforcement(self):
        """Jira ID:
        DAOS-6267: Verify that a container can be created with and enforces
                   a redundancy factor of 0.
        DAOS-6268: Verify container with RF 0 created, deleted and error
                   reported as target fails.
        DAOS-6269: Verify container with RF 2 enforces object creation limits.
        Description:
            Test step:
                (1)Create pool and container with redundant factor
                   Start background IO object write, and verify the
                   container redundant factor vs. different object_class
                   traffic.
                   Continue following steps for the positive testcases.
                (2)Check for container initial rf and health-status.
                (3)Rank rebuild start.
                (4)Check for container rf and health-status after the
                   rebuild started.
                (5)Check for container rf and health-status after the
                   rebuild completed.
                (6)Check for pool and container info after rebuild.
                (7)Verify container io object write if the container is
                   healthy.
        Use Cases:
            Verify container RF enforcement with different object class
            traffic, positive test of rebuild with server failure and read
            write io verification.

        :avocado: tags=all,full_regression
        :avocado: tags=container,rebuild
        :avocado: tags=container_rf
        :avocado: tags=cont_rf_oclass_enforcement
        """
        self.execute_rebuild_test()
