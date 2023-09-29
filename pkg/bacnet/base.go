package bacnet

type NetworkPriority int

const (
	LifeSafetyMessage NetworkPriority = iota + 1
	CriticalEquipmentMessage
	UrgentMessage
	NormalMessage
)

type NetworkLayerMessageType int

const (
	WhoIsRouterToNetwork NetworkLayerMessageType = iota
	IAmRouterToNetwork
	ICouldBeRouterToNetwork
	RejectMessageToNetwork
	RouterBusyToNetwork
	RouterAvailableToNetwork
	InitRtTable
	InitRtTableAck
	EstablishConnectionToNetwork
	DisconnectConnectionToNetwork
	ChallengeRequest
	SecurityPayload
	SecurityResponse
	RequestKeyUpdate
	UpdateKeySet
	UpdateDistributionKey
	RequestMasterKey
	SetMasterKey
	WhatIsNetworkNumber
	NetworkNumberIs
)

type MaxSegments int

const (
	MaxSEG0  MaxSegments = 0
	MaxSEG2  MaxSegments = 0x10
	MaxSEG4  MaxSegments = 0x20
	MaxSEG8  MaxSegments = 0x30
	MaxSEG16 MaxSegments = 0x40
	MaxSEG32 MaxSegments = 0x50
	MaxSEG64 MaxSegments = 0x60
	MaxSEG65 MaxSegments = 0x70
)

type PduTypes int

const (
	PDUTypeConfirmedServiceRequest PduTypes = iota
	Server
	NegativeAck
	// TODO
	SegmentResponseAccepted
	MoreFollows
	SegmentedMessage                 PduTypes = 8
	PDUTypeUnconfirmedServiceRequest PduTypes = 0x10
	PDUTypeSimpleAck                 PduTypes = 0x20
	PDUTypeComplexAck                PduTypes = 0x30
	PDUTypeSegmentAck                PduTypes = 0x40
	PDUTypeError                     PduTypes = 0x50
	PDUTypeReject                    PduTypes = 0x60
	PDUTypeAbort                     PduTypes = 0x70
	PDUTypeMask                      PduTypes = 0xF0
)

type NpduControl int

const (
	PriorityNormalMessage NpduControl = iota
	PriorityUrgentMessage
	PriorityCriticalMessage
	PriorityLifeSafetyMessage
	ExpectingReply
	SourceSpecified      NpduControl = 8
	DestinationSpecified NpduControl = 32
	NetworkLayerMessage  NpduControl = 128
)

type NetworkMessageType int

const (
	NetworkMessageWhoIsRouterToNetwork NetworkMessageType = iota
	NetworkMessageIAmRouterToNetwork
	NetworkMessageICouldBeRouterToNetwork
	NetworkMessageRejectMessageToNetwork
	NetworkMessageRouterBusyToNetwork
	NetworkMessageRouterAvailableToNetwork
	NetworkMessageInitRtTable
	NetworkMessageInitRtTableAck
	NetworkMessageEstablishConnectionToNetwork
	NetworkMessageDisconnectConnectionToNetwork
	NetworkMessageChallengeRequest
	NetworkMessageSecurityPayload
	NetworkMessageSecurityResponse
	NetworkMessageRequestKeyUpdate
	NetworkMessageUpdateKeySet
	NetworkMessageUpdateDistributionKey
	NetworkMessageRequestMasterKey
	NetworkMessageSetMasterKey
	NetworkMessageWhatIsNetworkNumber
	NetworkMessageNetworkNumberIs
)

// ErrorClassEnum represents the error classes as constants.
type ErrorClassEnum int

const (
	Device ErrorClassEnum = iota
	Object
	Property
	Resources
	Security
	Services
	VT
	Communication
)

// Error Code Enum represents the error codes as constants.
type ErrorCodeEnum int

const (
	Other ErrorCodeEnum = iota
	AuthenticationFailed
	ConfigurationInProgress
	DeviceBusy
	DynamicCreationNotSupported
	FileAccessDenied
	IncompatibleSecurityLevels
	InconsistentParameters
	InconsistentSelectionCriterion
	InvalidDataType
	InvalidFileAccessMethod
	InvalidFileStartPosition
	InvalidOperatorName
	InvalidParameterDataType
	InvalidTimestamp
	KeyGenerationError
	MissingRequiredParameter
	NoObjectsOfSpecifiedType
	NoSpaceForObject
	NoSpaceToAddListElement
	NoSpaceToWriteProperty
	NoVTSessionsAvailable
	PropertyIsNotAList
	ObjectDeletionNotPermitted
	ObjectIdentifierAlreadyExists
	OperationalProblem
	PasswordFailure
	ReadAccessDenied
	SecurityNotSupported
	ServiceRequestDenied
	Timeout
	UnknownObject
	UnknownProperty
	UnknownVTClass
	UnknownVTSession
	UnsupportedObjectType
	ValueOutOfRange
	VTSessionAlreadyClosed
	VTSessionTerminationFailure
	WriteAccessDenied
	CharacterSetNotSupported
	InvalidArrayIndex
	COVSubscriptionFailed
	NotCOVProperty
	OptionalFunctionalityNotSupported
	InvalidConfigurationData
	DataTypeNotSupported
	DuplicateName
	DuplicateObjectID
	PropertyIsNotAnArray
	AbortBufferOverflow
	AbortInvalidAPDUInThisState
	AbortPreemptedbyHigherPriorityTask
	AbortSegmentationNotSupported
	AbortProprietary
	AbortOther
	InvalidTag
	NetworkDown
	RejectBufferOverflow
	RejectInconsistentParameters
	RejectInvalidParameterDataType
	RejectInvalidTag
	RejectMissingRequiredParameter
	RejectParameterOutOfRange
	RejectTooManyArguments
	RejectUndefinedEnumeration
	RejectUnrecognizedService
	RejectProprietary
	RejectOther
	UnknownDevice
	UnknownRoute
	ValueNotInitialized
	InvalidEventState
	NoAlarmConfigured
	LogBufferFull
	LoggedValuePurged
	NoPropertySpecified
	NotConfiguredForTriggeredLogging
	UnknownSubscription
	ParameterOutOfRange
	ListElementNotFound
	Busy
	CommunicationDisabled
	Success
	AccessDenied
	BadDestinationAddress
	BadDestinationDeviceID
	BadSignature
	BadSourceAddress
	BadTimestamp
	CannotUseKey
	CannotVerifyMessageID
	CorrectKeyRevision
	DestinationDeviceIDRequired
	DuplicateMessage
	EncryptionNotConfigured
	EncryptionRequired
	IncorrectKey
	InvalidKeyData
	KeyUpdateInProgress
	MalformedMessage
	NotKeyServer
	SecurityNotConfigured
	SourceSecurityRequired
	TooManyKeys
	UnknownAuthenticationType
	UnknownKey
	UnknownKeyRevision
	UnknownSourceMessage
	NotRouterToDNET
	RouterBusy
	UnknownNetworkMessage
	MessageTooLong
	SecurityError
	AddressingError
	WriteBDTFailed
	ReadBDTFailed
	RegisterForeignDeviceFailed
	ReadFDTFailed
	DeleteFDTEntryFailed
	DistributeBroadcastFailed
	UnknownFileSize
	AbortAPDUTooLong
	AbortApplicationExceededReplyTime
	AbortOutOfResources
	AbortTSMTimeout
	AbortWindowSizeOutOfRange
	FileFull
	InconsistentConfiguration
	InconsistentObjectType
	InternalError
	NotConfigured
	OutOfMemory
	ValueTooLong
	AbortInsufficientSecurity
	AbortSecurityError
	DuplicateEntry
	InvalidValueInThisState
)
