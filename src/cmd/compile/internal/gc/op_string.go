// Code generated by "stringer -type=Op -trimprefix=O"; DO NOT EDIT.

package gc

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OXXX-0]
	_ = x[ONAME-1]
	_ = x[ONONAME-2]
	_ = x[OTYPE-3]
	_ = x[OPACK-4]
	_ = x[OLITERAL-5]
	_ = x[OADD-6]
	_ = x[OSUB-7]
	_ = x[OOR-8]
	_ = x[OXOR-9]
	_ = x[OADDSTR-10]
	_ = x[OADDR-11]
	_ = x[OANDAND-12]
	_ = x[OAPPEND-13]
	_ = x[OBYTES2STR-14]
	_ = x[OBYTES2STRTMP-15]
	_ = x[ORUNES2STR-16]
	_ = x[OSTR2BYTES-17]
	_ = x[OSTR2BYTESTMP-18]
	_ = x[OSTR2RUNES-19]
	_ = x[OAS-20]
	_ = x[OAS2-21]
	_ = x[OAS2DOTTYPE-22]
	_ = x[OAS2FUNC-23]
	_ = x[OAS2MAPR-24]
	_ = x[OAS2RECV-25]
	_ = x[OASOP-26]
	_ = x[OCALL-27]
	_ = x[OCALLFUNC-28]
	_ = x[OCALLMETH-29]
	_ = x[OCALLINTER-30]
	_ = x[OCALLPART-31]
	_ = x[OCAP-32]
	_ = x[OCLOSE-33]
	_ = x[OCLOSURE-34]
	_ = x[OCOMPLIT-35]
	_ = x[OMAPLIT-36]
	_ = x[OSTRUCTLIT-37]
	_ = x[OARRAYLIT-38]
	_ = x[OSLICELIT-39]
	_ = x[OPTRLIT-40]
	_ = x[OCONV-41]
	_ = x[OCONVIFACE-42]
	_ = x[OCONVNOP-43]
	_ = x[OCOPY-44]
	_ = x[ODCL-45]
	_ = x[ODCLFUNC-46]
	_ = x[ODCLFIELD-47]
	_ = x[ODCLCONST-48]
	_ = x[ODCLTYPE-49]
	_ = x[ODELETE-50]
	_ = x[ODOT-51]
	_ = x[ODOTPTR-52]
	_ = x[ODOTMETH-53]
	_ = x[ODOTINTER-54]
	_ = x[OXDOT-55]
	_ = x[ODOTTYPE-56]
	_ = x[ODOTTYPE2-57]
	_ = x[OEQ-58]
	_ = x[ONE-59]
	_ = x[OLT-60]
	_ = x[OLE-61]
	_ = x[OGE-62]
	_ = x[OGT-63]
	_ = x[ODEREF-64]
	_ = x[OINDEX-65]
	_ = x[OINDEXMAP-66]
	_ = x[OKEY-67]
	_ = x[OSTRUCTKEY-68]
	_ = x[OLEN-69]
	_ = x[OMAKE-70]
	_ = x[OMAKECHAN-71]
	_ = x[OMAKEMAP-72]
	_ = x[OMAKESLICE-73]
	_ = x[OMAKESLICECOPY-74]
	_ = x[OMUL-75]
	_ = x[ODIV-76]
	_ = x[OMOD-77]
	_ = x[OLSH-78]
	_ = x[ORSH-79]
	_ = x[OAND-80]
	_ = x[OANDNOT-81]
	_ = x[ONEW-82]
	_ = x[ONEWOBJ-83]
	_ = x[ONOT-84]
	_ = x[OBITNOT-85]
	_ = x[OPLUS-86]
	_ = x[ONEG-87]
	_ = x[OOROR-88]
	_ = x[OPANIC-89]
	_ = x[OPRINT-90]
	_ = x[OPRINTN-91]
	_ = x[OPAREN-92]
	_ = x[OSEND-93]
	_ = x[OSLICE-94]
	_ = x[OSLICEARR-95]
	_ = x[OSLICESTR-96]
	_ = x[OSLICE3-97]
	_ = x[OSLICE3ARR-98]
	_ = x[OSLICEHEADER-99]
	_ = x[ORECOVER-100]
	_ = x[ORECV-101]
	_ = x[ORUNESTR-102]
	_ = x[OSELRECV-103]
	_ = x[OSELRECV2-104]
	_ = x[OIOTA-105]
	_ = x[OREAL-106]
	_ = x[OIMAG-107]
	_ = x[OCOMPLEX-108]
	_ = x[OALIGNOF-109]
	_ = x[OOFFSETOF-110]
	_ = x[OSIZEOF-111]
	_ = x[OBLOCK-112]
	_ = x[OBREAK-113]
	_ = x[OCASE-114]
	_ = x[OCONTINUE-115]
	_ = x[ODEFER-116]
	_ = x[OEMPTY-117]
	_ = x[OFALL-118]
	_ = x[OFOR-119]
	_ = x[OFORUNTIL-120]
	_ = x[OGOTO-121]
	_ = x[OIF-122]
	_ = x[OLABEL-123]
	_ = x[OGO-124]
	_ = x[OGORDER-125]
	_ = x[ORANGE-126]
	_ = x[ORETURN-127]
	_ = x[OSELECT-128]
	_ = x[OSWITCH-129]
	_ = x[OTYPESW-130]
	_ = x[OTCHAN-131]
	_ = x[OTMAP-132]
	_ = x[OTSTRUCT-133]
	_ = x[OTINTER-134]
	_ = x[OTFUNC-135]
	_ = x[OTARRAY-136]
	_ = x[ODDD-137]
	_ = x[OINLCALL-138]
	_ = x[OEFACE-139]
	_ = x[OITAB-140]
	_ = x[OIDATA-141]
	_ = x[OSPTR-142]
	_ = x[OCLOSUREVAR-143]
	_ = x[OCFUNC-144]
	_ = x[OCHECKNIL-145]
	_ = x[OVARDEF-146]
	_ = x[OVARKILL-147]
	_ = x[OVARLIVE-148]
	_ = x[ORESULT-149]
	_ = x[OINLMARK-150]
	_ = x[ORETJMP-151]
	_ = x[OGETG-152]
	_ = x[OEND-153]
}

const _Op_name = "XXXNAMENONAMETYPEPACKLITERALADDSUBORXORADDSTRADDRANDANDAPPENDBYTES2STRBYTES2STRTMPRUNES2STRSTR2BYTESSTR2BYTESTMPSTR2RUNESASAS2AS2DOTTYPEAS2FUNCAS2MAPRAS2RECVASOPCALLCALLFUNCCALLMETHCALLINTERCALLPARTCAPCLOSECLOSURECOMPLITMAPLITSTRUCTLITARRAYLITSLICELITPTRLITCONVCONVIFACECONVNOPCOPYDCLDCLFUNCDCLFIELDDCLCONSTDCLTYPEDELETEDOTDOTPTRDOTMETHDOTINTERXDOTDOTTYPEDOTTYPE2EQNELTLEGEGTDEREFINDEXINDEXMAPKEYSTRUCTKEYLENMAKEMAKECHANMAKEMAPMAKESLICEMAKESLICECOPYMULDIVMODLSHRSHANDANDNOTNEWNEWOBJNOTBITNOTPLUSNEGORORPANICPRINTPRINTNPARENSENDSLICESLICEARRSLICESTRSLICE3SLICE3ARRSLICEHEADERRECOVERRECVRUNESTRSELRECVSELRECV2IOTAREALIMAGCOMPLEXALIGNOFOFFSETOFSIZEOFBLOCKBREAKCASECONTINUEDEFEREMPTYFALLFORFORUNTILGOTOIFLABELGOGORDERRANGERETURNSELECTSWITCHTYPESWTCHANTMAPTSTRUCTTINTERTFUNCTARRAYDDDINLCALLEFACEITABIDATASPTRCLOSUREVARCFUNCCHECKNILVARDEFVARKILLVARLIVERESULTINLMARKRETJMPGETGEND"

var _Op_index = [...]uint16{0, 3, 7, 13, 17, 21, 28, 31, 34, 36, 39, 45, 49, 55, 61, 70, 82, 91, 100, 112, 121, 123, 126, 136, 143, 150, 157, 161, 165, 173, 181, 190, 198, 201, 206, 213, 220, 226, 235, 243, 251, 257, 261, 270, 277, 281, 284, 291, 299, 307, 314, 320, 323, 329, 336, 344, 348, 355, 363, 365, 367, 369, 371, 373, 375, 380, 385, 393, 396, 405, 408, 412, 420, 427, 436, 449, 452, 455, 458, 461, 464, 467, 473, 476, 482, 485, 491, 495, 498, 502, 507, 512, 518, 523, 527, 532, 540, 548, 554, 563, 574, 581, 585, 592, 599, 607, 611, 615, 619, 626, 633, 641, 647, 652, 657, 661, 669, 674, 679, 683, 686, 694, 698, 700, 705, 707, 713, 718, 724, 730, 736, 742, 747, 751, 758, 764, 769, 775, 778, 785, 790, 794, 799, 803, 813, 818, 826, 832, 839, 846, 852, 859, 865, 869, 872}

func (i Op) String() string {
	if i >= Op(len(_Op_index)-1) {
		return "Op(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Op_name[_Op_index[i]:_Op_index[i+1]]
}
