#!/usr/bin/python
'''
  (C) Copyright 2020-2021 Intel Corporation.

  SPDX-License-Identifier: BSD-2-Clause-Patent
'''
from ior_test_base import IorTestBase


class EcodIor(IorTestBase):
    # pylint: disable=too-many-ancestors
    """EC IOR class to run smoke tests.

    Test Class Description: To validate Erasure code object type classes.

    :avocado: recursive
    """

    def test_ec(self):
        """Jira ID: DAOS-5812.

        Test Description:
            Test Erasure code object with IOR.
        Use Case:
            Create the medium size of pool and run IOR with supported EC object
            type class for sanity purpose.

        :avocado: tags=all,pr,daily_regression
        :avocado: tags=hw,large
        :avocado: tags=ec,ec_smoke,mpich
        :avocado: tags=ec_ior
        """
        obj_class = self.params.get("dfs_oclass", '/run/ior/objectclass/*')

        for oclass in obj_class:
            self.ior_cmd.dfs_oclass.update(oclass)
            self.ior_cmd.dfs_dir_oclass.update(oclass)
            self.add_pool_with_params(self.get_pool_params())
            self.pool.create()
            self.create_cont()
            self.ior_cmd.set_daos_params(
                self.server_group, self.pool, self.container.uuid)
            self.run_ior_with_pool(create_pool=False)
