#!/usr/bin/python3
"""
(C) Copyright 2021 Intel Corporation.

SPDX-License-Identifier: BSD-2-Clause-Patent
"""
import threading

from scrubber_utils import ScrubberUtils
from ior_test_base import IorTestBase


class TestWithScrubber(IorTestBase):
    """Test with scrubber enabled.

    :avocado: recursive
    """

    def __init__(self, *args, **kwargs):
        """Initialize a Test object."""
        super().__init__(*args, **kwargs)
        self.scrubber = None

    def setUp(self):
        """Set up each test case."""
        super().setUp()
        self.scrubber = ScrubberUtils(
            self.get_dmg_command(), self.server_managers[0].hosts)
        self.daos_cmd = self.get_daos_command()

    def create_pool_cont_with_scrubber(self, pool_prop=None, cont_prop=None):
        self.add_pool()
        self.add_container(pool=self.pool, create=False)
        if pool_prop is None:
            pool_prop = "scrub:continuous,scrub-freq:1,scrub-cred:10"
        if cont_prop is None:
            cont_prop = "cksum:crc16"
        x = pool_prop.split(",")
        for prop_val in x:
            if prop_val is not None:
                value = prop_val.split(":")
                self.pool.set_property(value[0], value[1])
        self.container.properties.value = cont_prop
        self.container.create()

    def run_ior_and_check_scruber_status(self, fail_on_warning=True):
        self.scrubber.get_csum_total_metrics()
        # Add a thread for these IOR arguments
        process = threading.Thread(target=self.run_ior_with_pool,
                                   kwargs={"create_pool": True,
                                           "create_cont": False,
                                           "fail_on_warning": fail_on_warning})
        # Launch the IOR thread
        process.start()
        # Wait for the thread to finish
        process.join()
        self.scrubber.get_csum_total_metrics()
