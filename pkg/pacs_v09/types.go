// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v09

// May be one of RADM, RPIN, FXDR, DISP, PUOR, SCOR
type DocumentType3Code string

// May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT, PUOR
type DocumentType6Code string

// Must match the pattern [0-9]{2}
type Exact2NumericText string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalCashClearingSystem1Code string

// Must be at least 1 items long
type ExternalCategoryPurpose1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalCreditorAgentInstruction1Code string

// Must be at least 1 items long
type ExternalDiscountAmountType1Code string

// Must be at least 1 items long
type ExternalDocumentLineType1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalGarnishmentType1Code string

// Must be at least 1 items long
type ExternalLocalInstrument1Code string

// Must be at least 1 items long
type ExternalMandateSetupReason1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

// Must be at least 1 items long
type ExternalProxyAccountType1Code string

// Must be at least 1 items long
type ExternalPurpose1Code string

// Must be at least 1 items long
type ExternalServiceLevel1Code string

// Must be at least 1 items long
type ExternalTaxAmountType1Code string

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA, FRTN
type Frequency6Code string

// May be one of PHOA, TELA
type Instruction4Code string

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

// May be one of HIGH, NORM
type Priority2Code string

// May be one of URGT, HIGH, NORM
type Priority3Code string

// May be one of CRED, DEBT, BOTH
type RegulatoryReportingType1Code string

// May be one of FAXI, EDIC, URID, EMAL, POST, SMSM
type RemittanceLocationMethod2Code string

// May be one of INDA, INGA, COVE, CLRG
type SettlementMethod1Code string

// May be one of MM01, MM02, MM03, MM04, MM05, MM06, MM07, MM08, MM09, MM10, MM11, MM12, QTR1, QTR2, QTR3, QTR4, HLF1, HLF2
type TaxRecordPeriod1Code string

// May be one of DEBT, CRED, SHAR, SLEV
type ChargeBearerType1Code string

// May be one of RTGS, RTNS, MPNS, BOOK
type ClearingChannel2Code string
