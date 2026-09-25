// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package stop

import (
	"fmt"
	"syscall"
)

type Errno syscall.Errno

func (errno Errno) String() string {
	if errno == 0 {
		return ""
	}
	if name, ok := errnoNames[errno]; ok {
		return name
	}
	return fmt.Sprintf("errno(%d)", uint32(errno))
}

// The instance runs Linux, so the errno in a stop code carries the Linux value,
// not the value of the host.
const (
	ErrnoE2BIG           Errno = 0x7
	ErrnoEACCES          Errno = 0xd
	ErrnoEADDRINUSE      Errno = 0x62
	ErrnoEADDRNOTAVAIL   Errno = 0x63
	ErrnoEADV            Errno = 0x44
	ErrnoEAFNOSUPPORT    Errno = 0x61
	ErrnoEAGAIN          Errno = 0xb
	ErrnoEALREADY        Errno = 0x72
	ErrnoEBADE           Errno = 0x34
	ErrnoEBADF           Errno = 0x9
	ErrnoEBADFD          Errno = 0x4d
	ErrnoEBADMSG         Errno = 0x4a
	ErrnoEBADR           Errno = 0x35
	ErrnoEBADRQC         Errno = 0x38
	ErrnoEBADSLT         Errno = 0x39
	ErrnoEBFONT          Errno = 0x3b
	ErrnoEBUSY           Errno = 0x10
	ErrnoECANCELED       Errno = 0x7d
	ErrnoECHILD          Errno = 0xa
	ErrnoECHRNG          Errno = 0x2c
	ErrnoECOMM           Errno = 0x46
	ErrnoECONNABORTED    Errno = 0x67
	ErrnoECONNREFUSED    Errno = 0x6f
	ErrnoECONNRESET      Errno = 0x68
	ErrnoEDEADLOCK       Errno = 0x23
	ErrnoEDESTADDRREQ    Errno = 0x59
	ErrnoEDOM            Errno = 0x21
	ErrnoEDOTDOT         Errno = 0x49
	ErrnoEDQUOT          Errno = 0x7a
	ErrnoEEXIST          Errno = 0x11
	ErrnoEFAULT          Errno = 0xe
	ErrnoEFBIG           Errno = 0x1b
	ErrnoEHOSTDOWN       Errno = 0x70
	ErrnoEHOSTUNREACH    Errno = 0x71
	ErrnoEHWPOISON       Errno = 0x85
	ErrnoEIDRM           Errno = 0x2b
	ErrnoEILSEQ          Errno = 0x54
	ErrnoEINPROGRESS     Errno = 0x73
	ErrnoEINTR           Errno = 0x4
	ErrnoEINVAL          Errno = 0x16
	ErrnoEIO             Errno = 0x5
	ErrnoEISCONN         Errno = 0x6a
	ErrnoEISDIR          Errno = 0x15
	ErrnoEISNAM          Errno = 0x78
	ErrnoEKEYEXPIRED     Errno = 0x7f
	ErrnoEKEYREJECTED    Errno = 0x81
	ErrnoEKEYREVOKED     Errno = 0x80
	ErrnoEL2HLT          Errno = 0x33
	ErrnoEL2NSYNC        Errno = 0x2d
	ErrnoEL3HLT          Errno = 0x2e
	ErrnoEL3RST          Errno = 0x2f
	ErrnoELIBACC         Errno = 0x4f
	ErrnoELIBBAD         Errno = 0x50
	ErrnoELIBEXEC        Errno = 0x53
	ErrnoELIBMAX         Errno = 0x52
	ErrnoELIBSCN         Errno = 0x51
	ErrnoELNRNG          Errno = 0x30
	ErrnoELOOP           Errno = 0x28
	ErrnoEMEDIUMTYPE     Errno = 0x7c
	ErrnoEMFILE          Errno = 0x18
	ErrnoEMLINK          Errno = 0x1f
	ErrnoEMSGSIZE        Errno = 0x5a
	ErrnoEMULTIHOP       Errno = 0x48
	ErrnoENAMETOOLONG    Errno = 0x24
	ErrnoENAVAIL         Errno = 0x77
	ErrnoENETDOWN        Errno = 0x64
	ErrnoENETRESET       Errno = 0x66
	ErrnoENETUNREACH     Errno = 0x65
	ErrnoENFILE          Errno = 0x17
	ErrnoENOANO          Errno = 0x37
	ErrnoENOBUFS         Errno = 0x69
	ErrnoENOCSI          Errno = 0x32
	ErrnoENODATA         Errno = 0x3d
	ErrnoENODEV          Errno = 0x13
	ErrnoENOENT          Errno = 0x2
	ErrnoENOEXEC         Errno = 0x8
	ErrnoENOKEY          Errno = 0x7e
	ErrnoENOLCK          Errno = 0x25
	ErrnoENOLINK         Errno = 0x43
	ErrnoENOMEDIUM       Errno = 0x7b
	ErrnoENOMEM          Errno = 0xc
	ErrnoENOMSG          Errno = 0x2a
	ErrnoENONET          Errno = 0x40
	ErrnoENOPKG          Errno = 0x41
	ErrnoENOPROTOOPT     Errno = 0x5c
	ErrnoENOSPC          Errno = 0x1c
	ErrnoENOSR           Errno = 0x3f
	ErrnoENOSTR          Errno = 0x3c
	ErrnoENOSYS          Errno = 0x26
	ErrnoENOTBLK         Errno = 0xf
	ErrnoENOTCONN        Errno = 0x6b
	ErrnoENOTDIR         Errno = 0x14
	ErrnoENOTEMPTY       Errno = 0x27
	ErrnoENOTNAM         Errno = 0x76
	ErrnoENOTRECOVERABLE Errno = 0x83
	ErrnoENOTSOCK        Errno = 0x58
	ErrnoENOTSUP         Errno = 0x5f
	ErrnoENOTTY          Errno = 0x19
	ErrnoENOTUNIQ        Errno = 0x4c
	ErrnoENXIO           Errno = 0x6
	ErrnoEOVERFLOW       Errno = 0x4b
	ErrnoEOWNERDEAD      Errno = 0x82
	ErrnoEPERM           Errno = 0x1
	ErrnoEPFNOSUPPORT    Errno = 0x60
	ErrnoEPIPE           Errno = 0x20
	ErrnoEPROTO          Errno = 0x47
	ErrnoEPROTONOSUPPORT Errno = 0x5d
	ErrnoEPROTOTYPE      Errno = 0x5b
	ErrnoERANGE          Errno = 0x22
	ErrnoEREMCHG         Errno = 0x4e
	ErrnoEREMOTE         Errno = 0x42
	ErrnoEREMOTEIO       Errno = 0x79
	ErrnoERESTART        Errno = 0x55
	ErrnoERFKILL         Errno = 0x84
	ErrnoEROFS           Errno = 0x1e
	ErrnoESHUTDOWN       Errno = 0x6c
	ErrnoESOCKTNOSUPPORT Errno = 0x5e
	ErrnoESPIPE          Errno = 0x1d
	ErrnoESRCH           Errno = 0x3
	ErrnoESRMNT          Errno = 0x45
	ErrnoESTALE          Errno = 0x74
	ErrnoESTRPIPE        Errno = 0x56
	ErrnoETIME           Errno = 0x3e
	ErrnoETIMEDOUT       Errno = 0x6e
	ErrnoETOOMANYREFS    Errno = 0x6d
	ErrnoETXTBSY         Errno = 0x1a
	ErrnoEUCLEAN         Errno = 0x75
	ErrnoEUNATCH         Errno = 0x31
	ErrnoEUSERS          Errno = 0x57
	ErrnoEXDEV           Errno = 0x12
	ErrnoEXFULL          Errno = 0x36
)

// errnoNames converts Linux errno values to their string representation.
var errnoNames = map[Errno]string{
	ErrnoE2BIG:           "E2BIG",
	ErrnoEACCES:          "EACCES",
	ErrnoEADDRINUSE:      "EADDRINUSE",
	ErrnoEADDRNOTAVAIL:   "EADDRNOTAVAIL",
	ErrnoEADV:            "EADV",
	ErrnoEAFNOSUPPORT:    "EAFNOSUPPORT",
	ErrnoEAGAIN:          "EAGAIN",
	ErrnoEALREADY:        "EALREADY",
	ErrnoEBADE:           "EBADE",
	ErrnoEBADF:           "EBADF",
	ErrnoEBADFD:          "EBADFD",
	ErrnoEBADMSG:         "EBADMSG",
	ErrnoEBADR:           "EBADR",
	ErrnoEBADRQC:         "EBADRQC",
	ErrnoEBADSLT:         "EBADSLT",
	ErrnoEBFONT:          "EBFONT",
	ErrnoEBUSY:           "EBUSY",
	ErrnoECANCELED:       "ECANCELED",
	ErrnoECHILD:          "ECHILD",
	ErrnoECHRNG:          "ECHRNG",
	ErrnoECOMM:           "ECOMM",
	ErrnoECONNABORTED:    "ECONNABORTED",
	ErrnoECONNREFUSED:    "ECONNREFUSED",
	ErrnoECONNRESET:      "ECONNRESET",
	ErrnoEDEADLOCK:       "EDEADLOCK",
	ErrnoEDESTADDRREQ:    "EDESTADDRREQ",
	ErrnoEDOM:            "EDOM",
	ErrnoEDOTDOT:         "EDOTDOT",
	ErrnoEDQUOT:          "EDQUOT",
	ErrnoEEXIST:          "EEXIST",
	ErrnoEFAULT:          "EFAULT",
	ErrnoEFBIG:           "EFBIG",
	ErrnoEHOSTDOWN:       "EHOSTDOWN",
	ErrnoEHOSTUNREACH:    "EHOSTUNREACH",
	ErrnoEHWPOISON:       "EHWPOISON",
	ErrnoEIDRM:           "EIDRM",
	ErrnoEILSEQ:          "EILSEQ",
	ErrnoEINPROGRESS:     "EINPROGRESS",
	ErrnoEINTR:           "EINTR",
	ErrnoEINVAL:          "EINVAL",
	ErrnoEIO:             "EIO",
	ErrnoEISCONN:         "EISCONN",
	ErrnoEISDIR:          "EISDIR",
	ErrnoEISNAM:          "EISNAM",
	ErrnoEKEYEXPIRED:     "EKEYEXPIRED",
	ErrnoEKEYREJECTED:    "EKEYREJECTED",
	ErrnoEKEYREVOKED:     "EKEYREVOKED",
	ErrnoEL2HLT:          "EL2HLT",
	ErrnoEL2NSYNC:        "EL2NSYNC",
	ErrnoEL3HLT:          "EL3HLT",
	ErrnoEL3RST:          "EL3RST",
	ErrnoELIBACC:         "ELIBACC",
	ErrnoELIBBAD:         "ELIBBAD",
	ErrnoELIBEXEC:        "ELIBEXEC",
	ErrnoELIBMAX:         "ELIBMAX",
	ErrnoELIBSCN:         "ELIBSCN",
	ErrnoELNRNG:          "ELNRNG",
	ErrnoELOOP:           "ELOOP",
	ErrnoEMEDIUMTYPE:     "EMEDIUMTYPE",
	ErrnoEMFILE:          "EMFILE",
	ErrnoEMLINK:          "EMLINK",
	ErrnoEMSGSIZE:        "EMSGSIZE",
	ErrnoEMULTIHOP:       "EMULTIHOP",
	ErrnoENAMETOOLONG:    "ENAMETOOLONG",
	ErrnoENAVAIL:         "ENAVAIL",
	ErrnoENETDOWN:        "ENETDOWN",
	ErrnoENETRESET:       "ENETRESET",
	ErrnoENETUNREACH:     "ENETUNREACH",
	ErrnoENFILE:          "ENFILE",
	ErrnoENOANO:          "ENOANO",
	ErrnoENOBUFS:         "ENOBUFS",
	ErrnoENOCSI:          "ENOCSI",
	ErrnoENODATA:         "ENODATA",
	ErrnoENODEV:          "ENODEV",
	ErrnoENOENT:          "ENOENT",
	ErrnoENOEXEC:         "ENOEXEC",
	ErrnoENOKEY:          "ENOKEY",
	ErrnoENOLCK:          "ENOLCK",
	ErrnoENOLINK:         "ENOLINK",
	ErrnoENOMEDIUM:       "ENOMEDIUM",
	ErrnoENOMEM:          "ENOMEM",
	ErrnoENOMSG:          "ENOMSG",
	ErrnoENONET:          "ENONET",
	ErrnoENOPKG:          "ENOPKG",
	ErrnoENOPROTOOPT:     "ENOPROTOOPT",
	ErrnoENOSPC:          "ENOSPC",
	ErrnoENOSR:           "ENOSR",
	ErrnoENOSTR:          "ENOSTR",
	ErrnoENOSYS:          "ENOSYS",
	ErrnoENOTBLK:         "ENOTBLK",
	ErrnoENOTCONN:        "ENOTCONN",
	ErrnoENOTDIR:         "ENOTDIR",
	ErrnoENOTEMPTY:       "ENOTEMPTY",
	ErrnoENOTNAM:         "ENOTNAM",
	ErrnoENOTRECOVERABLE: "ENOTRECOVERABLE",
	ErrnoENOTSOCK:        "ENOTSOCK",
	ErrnoENOTSUP:         "ENOTSUP",
	ErrnoENOTTY:          "ENOTTY",
	ErrnoENOTUNIQ:        "ENOTUNIQ",
	ErrnoENXIO:           "ENXIO",
	ErrnoEOVERFLOW:       "EOVERFLOW",
	ErrnoEOWNERDEAD:      "EOWNERDEAD",
	ErrnoEPERM:           "EPERM",
	ErrnoEPFNOSUPPORT:    "EPFNOSUPPORT",
	ErrnoEPIPE:           "EPIPE",
	ErrnoEPROTO:          "EPROTO",
	ErrnoEPROTONOSUPPORT: "EPROTONOSUPPORT",
	ErrnoEPROTOTYPE:      "EPROTOTYPE",
	ErrnoERANGE:          "ERANGE",
	ErrnoEREMCHG:         "EREMCHG",
	ErrnoEREMOTE:         "EREMOTE",
	ErrnoEREMOTEIO:       "EREMOTEIO",
	ErrnoERESTART:        "ERESTART",
	ErrnoERFKILL:         "ERFKILL",
	ErrnoEROFS:           "EROFS",
	ErrnoESHUTDOWN:       "ESHUTDOWN",
	ErrnoESOCKTNOSUPPORT: "ESOCKTNOSUPPORT",
	ErrnoESPIPE:          "ESPIPE",
	ErrnoESRCH:           "ESRCH",
	ErrnoESRMNT:          "ESRMNT",
	ErrnoESTALE:          "ESTALE",
	ErrnoESTRPIPE:        "ESTRPIPE",
	ErrnoETIME:           "ETIME",
	ErrnoETIMEDOUT:       "ETIMEDOUT",
	ErrnoETOOMANYREFS:    "ETOOMANYREFS",
	ErrnoETXTBSY:         "ETXTBSY",
	ErrnoEUCLEAN:         "EUCLEAN",
	ErrnoEUNATCH:         "EUNATCH",
	ErrnoEUSERS:          "EUSERS",
	ErrnoEXDEV:           "EXDEV",
	ErrnoEXFULL:          "EXFULL",
}
