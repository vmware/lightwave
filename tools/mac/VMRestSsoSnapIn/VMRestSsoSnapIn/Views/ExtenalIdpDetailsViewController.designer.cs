/*
 * Copyright © 2012-2016 VMware, Inc.  All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the “License”); you may not
 * use this file except in compliance with the License.  You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an “AS IS” BASIS, without
 * warranties or conditions of any kind, EITHER EXPRESS OR IMPLIED.  See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

// WARNING
//
// This file has been generated automatically by Xamarin Studio Business to store outlets and
// actions made in the UI designer. If it is removed, they will be lost.
// Manual changes to this file may not be handled correctly.
//
using Foundation;
using System.CodeDom.Compiler;

namespace RestSsoAdminSnapIn
{
	[Register ("ExtenalIdpDetailsViewController")]
	partial class ExtenalIdpDetailsViewController
	{
		[Outlet]
		AppKit.NSButton BtnJit { get; set; }

		[Outlet]
		AppKit.NSButton BtnViewCertificate { get; set; }

		[Outlet]
		AppKit.NSTableView CertificateTableView { get; set; }

		[Outlet]
		AppKit.NSButton EditButton { get; set; }

		[Outlet]
		AppKit.NSTableView NameFormatTableView { get; set; }

		[Outlet]
		AppKit.NSTableView SloTableView { get; set; }

		[Outlet]
		AppKit.NSTableView SsoTableView { get; set; }

		[Outlet]
		AppKit.NSTableView SubjectFormatTableView { get; set; }

		[Outlet]
		AppKit.NSTextField TxtAlias { get; set; }

		[Outlet]
		AppKit.NSTextField TxtEntityName { get; set; }
		
		void ReleaseDesignerOutlets ()
		{
			if (BtnJit != null) {
				BtnJit.Dispose ();
				BtnJit = null;
			}

			if (BtnViewCertificate != null) {
				BtnViewCertificate.Dispose ();
				BtnViewCertificate = null;
			}

			if (CertificateTableView != null) {
				CertificateTableView.Dispose ();
				CertificateTableView = null;
			}

			if (NameFormatTableView != null) {
				NameFormatTableView.Dispose ();
				NameFormatTableView = null;
			}

			if (SloTableView != null) {
				SloTableView.Dispose ();
				SloTableView = null;
			}

			if (SsoTableView != null) {
				SsoTableView.Dispose ();
				SsoTableView = null;
			}

			if (SubjectFormatTableView != null) {
				SubjectFormatTableView.Dispose ();
				SubjectFormatTableView = null;
			}

			if (TxtAlias != null) {
				TxtAlias.Dispose ();
				TxtAlias = null;
			}

			if (TxtEntityName != null) {
				TxtEntityName.Dispose ();
				TxtEntityName = null;
			}

			if (EditButton != null) {
				EditButton.Dispose ();
				EditButton = null;
			}
		}
	}
}
