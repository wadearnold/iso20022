// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pacs_v11

type DocumentPacs00200111 struct {
	FIToFIPmtStsRpt FIToFIPaymentStatusReportV11 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.11 FIToFIPmtStsRpt"`
}
