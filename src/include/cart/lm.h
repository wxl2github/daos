/* Copyright (C) 2016-2017 Intel Corporation
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted for any purpose (including commercial purposes)
 * provided that the following conditions are met:
 *
 * 1. Redistributions of source code must retain the above copyright notice,
 *    this list of conditions, and the following disclaimer.
 *
 * 2. Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions, and the following disclaimer in the
 *    documentation and/or materials provided with the distribution.
 *
 * 3. In addition, redistributions of modified forms of the source or binary
 *    code must carry prominent notices stating that the original code was
 *    changed and the date of the change.
 *
 *  4. All publications or advertising materials mentioning features or use of
 *     this software are asked, but not required, to acknowledge that it was
 *     developed by Intel Corporation and credit the contributors.
 *
 * 5. Neither the name of Intel Corporation, nor the name of any Contributor
 *    may be used to endorse or promote products derived from this software
 *    without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER BE LIABLE FOR ANY
 * DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
 * ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
 * THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */
/**
 * CaRT (Collective and RPC Transport) fault tolerance module management APIs.
 */

#ifndef __CRT_LM_H__
#define __CRT_LM_H__

#if defined(__cplusplus)
extern "C" {
#endif

#include <pouch/common.h>

/**
 * Enable the fault tolerance module. This function should be called by every
 * rank in the same group, and should be called after crt_init(). This function
 * will:
 *	1) turn on cart context progress callbacks
 *	2) turn on RAS notification event handlers
 */
void
crt_lm_init(void);

/**
 * Finilize the fault tolerance module. This function should be called before
 * crt_finalize().
 */
void
crt_lm_finalize(void);

/**
 * Retrieve the PSR candidate list for \a tgt_grp.
 * \param tgt_grp [IN]		The remote group
 * \param psr_cand [OUT]	The PSR candidate list for \a tgt_grp. The first
 *				entry of psr_cand is the current PSR. The rest
 *				of the list are backup PSRs. User should call
 *				crt_rank_list_free() to free the memory after
 *				using it.
 * \return			0 on success, negative value on error.
 */
int
crt_lm_group_psr(crt_group_t *tgt_grp, crt_rank_list_t **psr_cand);

#if defined(__cplusplus)
}
#endif

#endif /* __CRT_LM_H__ */
