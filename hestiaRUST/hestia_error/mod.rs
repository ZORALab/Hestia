// Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
//
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
//                  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.
pub enum Error {
	// Common errors and operations
	Ok,
	BadDescriptor,
	BadExchange,
	BadExec,
	BadMount,
	BadPipe,
	BadRequest,
	BadStreamPipe,
	Cancelled,
	CleaningRequired,
	Deadlock,
	Expired,
	IllegalByteSequence,
	IllegalSeek,
	InvalidArgument,
	IsEmpty,
	MaxedExchange,
	MaxedQuota,
	MissingLock,
	NotEmpty,
	NotPermitted,
	NotPossible,
	NotPossibleByRFKill,
	NotRecoverable,
	OutOfRange,
	PermissionDenied,
	Timeout,
	TooManyRead,
	TooManyLoop,
	TooManyReferences,
	TooManyLink,
	TryAgain,
	Unsupported,
	WouldBlock,

	// lifecycle states
	Restart,
	Resume,
	Shutdown,
	Sleep,
	Stalled,
	Standby,

	// progress
	ProgressScheduled,
	ProgressAlreadyExecuting,
	ProgressExecuting,
	ProgressCompleted,

	// tri-tier inter-package communications
	LV1NotSync,
	LV1Paused,
	LV1Reset,

	LV2NotSync,
	LV2Paused,
	LV2Reset,

	LV3NotSync,
	LV3Paused,
	LV3Reset,

	// data (input/output parameters, type, etc)
	DataBad,
	DataEmpty,
	DataInvalid,
	DataIsUnique,
	DataMissing,
	DataNotUnique,
	DataOverflow,
	DataRemoved,
	DataTooLong,

	// entity (device, file, directory, etc)
	EntityBad,
	EntityBusy,
	EntityDead,
	EntityExists,
	EntityFaulty,
	EntityMissing,
	EntityMissingChild,
	EntityOutOfBuffer,
	EntityPoisoned,
	EntityReadOnly,
	EntityTooBig,
	EntityTooManyOpened,
	EntityUnattached,

	EntityIsDirectory,
	EntityIsFile,
	EntityIsLink,
	EntityIsSocket,

	EntityIsNotDirectory,
	EntityIsNotFile,
	EntityIsNotLink,
	EntityIsNotSocket,

	EntityRemoteChanged,
	EntityRemoteError,
	EntityRemoteIO,

	EntityMissingStreamableResources,
	EntityNotStreamable,
	EntityStreamable,

	EntityATypewriter,
	EntityNotATypewriter,

	EntityBadDescriptor,
	EntityFiletableOverflow,

	// key (cryptography)
	KeyBad,
	KeyDestroyed,
	KeyExpired,
	KeyMissing,
	KeyRejected,
	KeyRevoked,

	// library
	LibraryBad,
	LibraryCorrupted,
	LibraryExecFailed,
	LibraryMaxed,
	LibraryMissing,

	// network
	NetworkBad,
	NetworkBadAd,
	NetworkDown,
	NetworkNotConnected,
	NetworkReset,
	NetworkRFS,
	NetworkUnreachable,

	NetworkHostDown,
	NetworkHostUnreachable,
	NetworkSocketUnsupported,

	NetworkAddressInUse,
	NetworkAddressUnavailable,

	NetworkConnAborted,
	NetworkConnIsConnected,
	NetworkConnMissingDestAddress,
	NetworkConnMultihop,
	NetworkConnNotConnected,
	NetworkConnRefused,
	NetworkConnReset,

	NetworkPayloadBad,
	NetworkPayloadEmpty,
	NetworkPayloadMissing,
	NetworkPayloadTooLong,

	// protocol
	ProtocolAddressUnsupported,
	ProtocolBad,
	ProtocolFamilyUnsupported,
	ProtocolMissing,
	ProtocolUnsupported,

	// system (e.g. os, interactable system)
	SystemBadIO,
	SystemDeviceCrossLink,
	SystemInterruptCall,
	SystemInvalid,
	SystemMissingBlockDevice,
	SystemMissingDevice,
	SystemMissingIO,
	SystemMissingProcess,
	SystemOutOfDomain,
	SystemOutOfMemory,
	SystemOutOfSpace,
	SystemReadOnlyFilesystem,

	// user
	UserAccessBanned,
	UserAccessLocked,
	UserAccessNotVerified,
	UserAccessRejected,
	UserAccessRevoked,

	UserIDBad,
	UserIDExists,
	UserIDMissing,
	UserMFABad,
	UserMFAExpired,
	UserMFAMissing,
	UserPasswordBad,
	UserPasswordExpired,
	UserPasswordMissing,
	UserKeyfileBad,
	UserKeyfileExpired,
	UserKeyfileMissing,

	UserAccessTokenBad,
	UserAccessTokenExpired,
	UserAccessTokenMissing,
	UserAccessTokenRevoked,
}
