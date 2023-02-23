// Code generated by "stringer -type=FeatureID,Vendor"; DO NOT EDIT.

package cpuid

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ADX-1]
	_ = x[AESNI-2]
	_ = x[AMD3DNOW-3]
	_ = x[AMD3DNOWEXT-4]
	_ = x[AMXBF16-5]
	_ = x[AMXINT8-6]
	_ = x[AMXTILE-7]
	_ = x[AVX-8]
	_ = x[AVX2-9]
	_ = x[AVX512BF16-10]
	_ = x[AVX512BITALG-11]
	_ = x[AVX512BW-12]
	_ = x[AVX512CD-13]
	_ = x[AVX512DQ-14]
	_ = x[AVX512ER-15]
	_ = x[AVX512F-16]
	_ = x[AVX512FP16-17]
	_ = x[AVX512IFMA-18]
	_ = x[AVX512PF-19]
	_ = x[AVX512VBMI-20]
	_ = x[AVX512VBMI2-21]
	_ = x[AVX512VL-22]
	_ = x[AVX512VNNI-23]
	_ = x[AVX512VP2INTERSECT-24]
	_ = x[AVX512VPOPCNTDQ-25]
	_ = x[AVXSLOW-26]
	_ = x[AVXVNNI-27]
	_ = x[BMI1-28]
	_ = x[BMI2-29]
	_ = x[CETIBT-30]
	_ = x[CETSS-31]
	_ = x[CLDEMOTE-32]
	_ = x[CLMUL-33]
	_ = x[CLZERO-34]
	_ = x[CMOV-35]
	_ = x[CMPSB_SCADBS_SHORT-36]
	_ = x[CMPXCHG8-37]
	_ = x[CPBOOST-38]
	_ = x[CX16-39]
	_ = x[ENQCMD-40]
	_ = x[ERMS-41]
	_ = x[F16C-42]
	_ = x[FMA3-43]
	_ = x[FMA4-44]
	_ = x[FXSR-45]
	_ = x[FXSROPT-46]
	_ = x[GFNI-47]
	_ = x[HLE-48]
	_ = x[HRESET-49]
	_ = x[HTT-50]
	_ = x[HWA-51]
	_ = x[HYPERVISOR-52]
	_ = x[IBPB-53]
	_ = x[IBS-54]
	_ = x[IBSBRNTRGT-55]
	_ = x[IBSFETCHSAM-56]
	_ = x[IBSFFV-57]
	_ = x[IBSOPCNT-58]
	_ = x[IBSOPCNTEXT-59]
	_ = x[IBSOPSAM-60]
	_ = x[IBSRDWROPCNT-61]
	_ = x[IBSRIPINVALIDCHK-62]
	_ = x[IBS_PREVENTHOST-63]
	_ = x[INT_WBINVD-64]
	_ = x[INVLPGB-65]
	_ = x[LAHF-66]
	_ = x[LAM-67]
	_ = x[LBRVIRT-68]
	_ = x[LZCNT-69]
	_ = x[MCAOVERFLOW-70]
	_ = x[MCOMMIT-71]
	_ = x[MMX-72]
	_ = x[MMXEXT-73]
	_ = x[MOVBE-74]
	_ = x[MOVDIR64B-75]
	_ = x[MOVDIRI-76]
	_ = x[MOVSB_ZL-77]
	_ = x[MPX-78]
	_ = x[MSRIRC-79]
	_ = x[MSR_PAGEFLUSH-80]
	_ = x[NRIPS-81]
	_ = x[NX-82]
	_ = x[OSXSAVE-83]
	_ = x[PCONFIG-84]
	_ = x[POPCNT-85]
	_ = x[RDPRU-86]
	_ = x[RDRAND-87]
	_ = x[RDSEED-88]
	_ = x[RDTSCP-89]
	_ = x[RTM-90]
	_ = x[RTM_ALWAYS_ABORT-91]
	_ = x[SCE-92]
	_ = x[SERIALIZE-93]
	_ = x[SEV-94]
	_ = x[SEV_64BIT-95]
	_ = x[SEV_ALTERNATIVE-96]
	_ = x[SEV_DEBUGSWAP-97]
	_ = x[SEV_ES-98]
	_ = x[SEV_RESTRICTED-99]
	_ = x[SEV_SNP-100]
	_ = x[SGX-101]
	_ = x[SGXLC-102]
	_ = x[SHA-103]
	_ = x[SME-104]
	_ = x[SME_COHERENT-105]
	_ = x[SSE-106]
	_ = x[SSE2-107]
	_ = x[SSE3-108]
	_ = x[SSE4-109]
	_ = x[SSE42-110]
	_ = x[SSE4A-111]
	_ = x[SSSE3-112]
	_ = x[STIBP-113]
	_ = x[STOSB_SHORT-114]
	_ = x[SUCCOR-115]
	_ = x[SVM-116]
	_ = x[SVMDA-117]
	_ = x[SVMFBASID-118]
	_ = x[SVML-119]
	_ = x[SVMNP-120]
	_ = x[SVMPF-121]
	_ = x[SVMPFT-122]
	_ = x[TBM-123]
	_ = x[TME-124]
	_ = x[TSCRATEMSR-125]
	_ = x[TSXLDTRK-126]
	_ = x[VAES-127]
	_ = x[VMCBCLEAN-128]
	_ = x[VMPL-129]
	_ = x[VMSA_REGPROT-130]
	_ = x[VMX-131]
	_ = x[VPCLMULQDQ-132]
	_ = x[VTE-133]
	_ = x[WAITPKG-134]
	_ = x[WBNOINVD-135]
	_ = x[X87-136]
	_ = x[XGETBV1-137]
	_ = x[XOP-138]
	_ = x[XSAVE-139]
	_ = x[XSAVEC-140]
	_ = x[XSAVEOPT-141]
	_ = x[XSAVES-142]
	_ = x[AESARM-143]
	_ = x[ARMCPUID-144]
	_ = x[ASIMD-145]
	_ = x[ASIMDDP-146]
	_ = x[ASIMDHP-147]
	_ = x[ASIMDRDM-148]
	_ = x[ATOMICS-149]
	_ = x[CRC32-150]
	_ = x[DCPOP-151]
	_ = x[EVTSTRM-152]
	_ = x[FCMA-153]
	_ = x[FP-154]
	_ = x[FPHP-155]
	_ = x[GPA-156]
	_ = x[JSCVT-157]
	_ = x[LRCPC-158]
	_ = x[PMULL-159]
	_ = x[SHA1-160]
	_ = x[SHA2-161]
	_ = x[SHA3-162]
	_ = x[SHA512-163]
	_ = x[SM3-164]
	_ = x[SM4-165]
	_ = x[SVE-166]
	_ = x[lastID-167]
	_ = x[firstID-0]
}

const _FeatureID_name = "firstIDADXAESNIAMD3DNOWAMD3DNOWEXTAMXBF16AMXINT8AMXTILEAVXAVX2AVX512BF16AVX512BITALGAVX512BWAVX512CDAVX512DQAVX512ERAVX512FAVX512FP16AVX512IFMAAVX512PFAVX512VBMIAVX512VBMI2AVX512VLAVX512VNNIAVX512VP2INTERSECTAVX512VPOPCNTDQAVXSLOWAVXVNNIBMI1BMI2CETIBTCETSSCLDEMOTECLMULCLZEROCMOVCMPSB_SCADBS_SHORTCMPXCHG8CPBOOSTCX16ENQCMDERMSF16CFMA3FMA4FXSRFXSROPTGFNIHLEHRESETHTTHWAHYPERVISORIBPBIBSIBSBRNTRGTIBSFETCHSAMIBSFFVIBSOPCNTIBSOPCNTEXTIBSOPSAMIBSRDWROPCNTIBSRIPINVALIDCHKIBS_PREVENTHOSTINT_WBINVDINVLPGBLAHFLAMLBRVIRTLZCNTMCAOVERFLOWMCOMMITMMXMMXEXTMOVBEMOVDIR64BMOVDIRIMOVSB_ZLMPXMSRIRCMSR_PAGEFLUSHNRIPSNXOSXSAVEPCONFIGPOPCNTRDPRURDRANDRDSEEDRDTSCPRTMRTM_ALWAYS_ABORTSCESERIALIZESEVSEV_64BITSEV_ALTERNATIVESEV_DEBUGSWAPSEV_ESSEV_RESTRICTEDSEV_SNPSGXSGXLCSHASMESME_COHERENTSSESSE2SSE3SSE4SSE42SSE4ASSSE3STIBPSTOSB_SHORTSUCCORSVMSVMDASVMFBASIDSVMLSVMNPSVMPFSVMPFTTBMTMETSCRATEMSRTSXLDTRKVAESVMCBCLEANVMPLVMSA_REGPROTVMXVPCLMULQDQVTEWAITPKGWBNOINVDX87XGETBV1XOPXSAVEXSAVECXSAVEOPTXSAVESAESARMARMCPUIDASIMDASIMDDPASIMDHPASIMDRDMATOMICSCRC32DCPOPEVTSTRMFCMAFPFPHPGPAJSCVTLRCPCPMULLSHA1SHA2SHA3SHA512SM3SM4SVElastID"

var _FeatureID_index = [...]uint16{0, 7, 10, 15, 23, 34, 41, 48, 55, 58, 62, 72, 84, 92, 100, 108, 116, 123, 133, 143, 151, 161, 172, 180, 190, 208, 223, 230, 237, 241, 245, 251, 256, 264, 269, 275, 279, 297, 305, 312, 316, 322, 326, 330, 334, 338, 342, 349, 353, 356, 362, 365, 368, 378, 382, 385, 395, 406, 412, 420, 431, 439, 451, 467, 482, 492, 499, 503, 506, 513, 518, 529, 536, 539, 545, 550, 559, 566, 574, 577, 583, 596, 601, 603, 610, 617, 623, 628, 634, 640, 646, 649, 665, 668, 677, 680, 689, 704, 717, 723, 737, 744, 747, 752, 755, 758, 770, 773, 777, 781, 785, 790, 795, 800, 805, 816, 822, 825, 830, 839, 843, 848, 853, 859, 862, 865, 875, 883, 887, 896, 900, 912, 915, 925, 928, 935, 943, 946, 953, 956, 961, 967, 975, 981, 987, 995, 1000, 1007, 1014, 1022, 1029, 1034, 1039, 1046, 1050, 1052, 1056, 1059, 1064, 1069, 1074, 1078, 1082, 1086, 1092, 1095, 1098, 1101, 1107}

func (i FeatureID) String() string {
	if i < 0 || i >= FeatureID(len(_FeatureID_index)-1) {
		return "FeatureID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FeatureID_name[_FeatureID_index[i]:_FeatureID_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[VendorUnknown-0]
	_ = x[Intel-1]
	_ = x[AMD-2]
	_ = x[VIA-3]
	_ = x[Transmeta-4]
	_ = x[NSC-5]
	_ = x[KVM-6]
	_ = x[MSVM-7]
	_ = x[VMware-8]
	_ = x[XenHVM-9]
	_ = x[Bhyve-10]
	_ = x[Hygon-11]
	_ = x[SiS-12]
	_ = x[RDC-13]
	_ = x[Ampere-14]
	_ = x[ARM-15]
	_ = x[Broadcom-16]
	_ = x[Cavium-17]
	_ = x[DEC-18]
	_ = x[Fujitsu-19]
	_ = x[Infineon-20]
	_ = x[Motorola-21]
	_ = x[NVIDIA-22]
	_ = x[AMCC-23]
	_ = x[Qualcomm-24]
	_ = x[Marvell-25]
	_ = x[lastVendor-26]
}

const _Vendor_name = "VendorUnknownIntelAMDVIATransmetaNSCKVMMSVMVMwareXenHVMBhyveHygonSiSRDCAmpereARMBroadcomCaviumDECFujitsuInfineonMotorolaNVIDIAAMCCQualcommMarvelllastVendor"

var _Vendor_index = [...]uint8{0, 13, 18, 21, 24, 33, 36, 39, 43, 49, 55, 60, 65, 68, 71, 77, 80, 88, 94, 97, 104, 112, 120, 126, 130, 138, 145, 155}

func (i Vendor) String() string {
	if i < 0 || i >= Vendor(len(_Vendor_index)-1) {
		return "Vendor(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Vendor_name[_Vendor_index[i]:_Vendor_index[i+1]]
}