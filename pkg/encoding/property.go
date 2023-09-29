package encoding

type NetworkType int

const (
	Ethernet NetworkType = iota
	ARCnet
	MSTP
	PTP
	LonTalk
	IPV4
	Zigbee
	Virtual
	IPV6
	Serial
)

type ObjectType uint16

const (
	AnalogInput ObjectType = iota
	AnalogOutput
	AnalogValue
	BinaryInput
	BinaryOutput
	BinaryValue
	Calendar
	ObjectTypeCommand
	ObjectTypeDevice
	EventEnrollment
	File
	Group
	Loop
	MultiStateInput
	MultiStateOutput
	ObjectTypeNotificationClass
	Program
	Schedule
	Averaging
	MultiStateValue
	TrendLog
	LifeSafetyPoint
	LifeSafetyZone
	Accumulator
	PulseConverter
	EventLog
	GlobalGroup
	TrendLogMultiple
	LoadControl
	StructuredView
	AccessDoor
	Timer
	AccessCredential
	AccessPoint
	AccessRights
	AccessUser
	AccessZone
	CredentialDataInput
	NetworkSecurity
	BitStringValue
	CharacterStringValue
	DatePatternValue
	DateValue
	DateTimePatternValue
	DateTimeValue
	IntegerValue
	LargeAnalogValue
	OctetStringValue
	PositiveIntegerValue
	TimePatternValue
	TimeValue
	NotificationForwarder
	AlertEnrollment
	Channel
	LightingOutput
	BinaryLightingOutput
	NetworkPort
	ObjectTypeElevatorGroup
	Escalator
	Lift
	Staging
)

// PropertyIdentifier represents  property identifiers.
type PropertyIdentifier int

const (
	AckedTransitions                       PropertyIdentifier = 0
	AckRequired                            PropertyIdentifier = 1
	PropertyIdentifierAction               PropertyIdentifier = 2
	ActionText                             PropertyIdentifier = 3
	ActiveText                             PropertyIdentifier = 4
	ActiveVtSessions                       PropertyIdentifier = 5
	AlarmValue                             PropertyIdentifier = 6
	AlarmValues                            PropertyIdentifier = 7
	All                                    PropertyIdentifier = 8
	AllWritesSuccessful                    PropertyIdentifier = 9
	ApduSegmentTimeout                     PropertyIdentifier = 10
	ApduTimeout                            PropertyIdentifier = 11
	ApplicationSoftwareVersion             PropertyIdentifier = 12
	Archive                                PropertyIdentifier = 13
	Bias                                   PropertyIdentifier = 14
	ChangeOfStateCount                     PropertyIdentifier = 15
	ChangeOfStateTime                      PropertyIdentifier = 16
	NotificationClass                      PropertyIdentifier = 17
	Blank1                                 PropertyIdentifier = 18
	ControlledVariableReference            PropertyIdentifier = 19
	ControlledVariableUnits                PropertyIdentifier = 20
	ControlledVariableValue                PropertyIdentifier = 21
	CovIncrement                           PropertyIdentifier = 22
	DateList                               PropertyIdentifier = 23
	DaylightSavingsStatus                  PropertyIdentifier = 24
	Deadband                               PropertyIdentifier = 25
	DerivativeConstant                     PropertyIdentifier = 26
	DerivativeConstantUnits                PropertyIdentifier = 27
	Description                            PropertyIdentifier = 28
	DescriptionOfHalt                      PropertyIdentifier = 29
	DeviceAddressBinding                   PropertyIdentifier = 30
	DeviceType                             PropertyIdentifier = 31
	EffectivePeriod                        PropertyIdentifier = 32
	ElapsedActiveTime                      PropertyIdentifier = 33
	ErrorLimit                             PropertyIdentifier = 34
	EventEnable                            PropertyIdentifier = 35
	PropertyIdentifierEventState           PropertyIdentifier = 36
	PropertyIdentifierEventType            PropertyIdentifier = 37
	ExceptionSchedule                      PropertyIdentifier = 38
	FaultValues                            PropertyIdentifier = 39
	FeedbackValue                          PropertyIdentifier = 40
	PropertyIdentifierFileAccessMethod     PropertyIdentifier = 41
	FileSize                               PropertyIdentifier = 42
	FileType                               PropertyIdentifier = 43
	FirmwareRevision                       PropertyIdentifier = 44
	HighLimit                              PropertyIdentifier = 45
	InactiveText                           PropertyIdentifier = 46
	InProcess                              PropertyIdentifier = 47
	InstanceOf                             PropertyIdentifier = 48
	IntegralConstant                       PropertyIdentifier = 49
	IntegralConstantUnits                  PropertyIdentifier = 50
	IssueConfirmedNotifications            PropertyIdentifier = 51
	LimitEnable                            PropertyIdentifier = 52
	ListOfGroupMembers                     PropertyIdentifier = 53
	ListOfObjectPropertyReferences         PropertyIdentifier = 54
	ListOfSessionKeys                      PropertyIdentifier = 55
	LocalDate                              PropertyIdentifier = 56
	LocalTime                              PropertyIdentifier = 57
	Location                               PropertyIdentifier = 58
	LowLimit                               PropertyIdentifier = 59
	ManipulatedVariableReference           PropertyIdentifier = 60
	MaximumOutput                          PropertyIdentifier = 61
	MaxApduLengthAccepted                  PropertyIdentifier = 62
	MaxInfoFrames                          PropertyIdentifier = 63
	MaxMaster                              PropertyIdentifier = 64
	MaxPresValue                           PropertyIdentifier = 65
	MinimumOffTime                         PropertyIdentifier = 66
	MinimumOnTime                          PropertyIdentifier = 67
	MinimumOutput                          PropertyIdentifier = 68
	MinPresValue                           PropertyIdentifier = 69
	ModelName                              PropertyIdentifier = 70
	ModificationDate                       PropertyIdentifier = 71
	PropertyIdentifierNotifyType           PropertyIdentifier = 72
	NumberOfApduRetries                    PropertyIdentifier = 73
	NumberOfStates                         PropertyIdentifier = 74
	ObjectIdentifierPI                     PropertyIdentifier = 75
	ObjectList                             PropertyIdentifier = 76
	ObjectName                             PropertyIdentifier = 77
	ObjectPropertyReference                PropertyIdentifier = 78
	ObjectTypePI                           PropertyIdentifier = 79
	Optional                               PropertyIdentifier = 80
	OutOfService                           PropertyIdentifier = 81
	OutputUnits                            PropertyIdentifier = 82
	EventParameters                        PropertyIdentifier = 83
	PropertyIdentifierPolarity             PropertyIdentifier = 84
	PresentValue                           PropertyIdentifier = 85
	Priority                               PropertyIdentifier = 86
	PriorityArray                          PropertyIdentifier = 87
	PriorityForWriting                     PropertyIdentifier = 88
	ProcessIdentifier                      PropertyIdentifier = 89
	ProgramChange                          PropertyIdentifier = 90
	ProgramLocation                        PropertyIdentifier = 91
	PropertyIdentifierProgramState         PropertyIdentifier = 92
	ProportionalConstant                   PropertyIdentifier = 93
	ProportionalConstantUnits              PropertyIdentifier = 94
	ProtocolConformanceClass               PropertyIdentifier = 95
	ProtocolObjectTypesSupported           PropertyIdentifier = 96
	ProtocolServicesSupported              PropertyIdentifier = 97
	ProtocolVersion                        PropertyIdentifier = 98
	ReadOnly                               PropertyIdentifier = 99
	ReasonForHalt                          PropertyIdentifier = 100
	Recipient                              PropertyIdentifier = 101
	RecipientList                          PropertyIdentifier = 102
	PropertyIdentifierReliability          PropertyIdentifier = 103
	RelinquishDefault                      PropertyIdentifier = 104
	Required                               PropertyIdentifier = 105
	Resolution                             PropertyIdentifier = 106
	SegmentationSupported                  PropertyIdentifier = 107
	Setpoint                               PropertyIdentifier = 108
	SetpointReference                      PropertyIdentifier = 109
	StateText                              PropertyIdentifier = 110
	StatusFlags                            PropertyIdentifier = 111
	SystemStatus                           PropertyIdentifier = 112
	TimeDelay                              PropertyIdentifier = 113
	TimeOfActiveTimeReset                  PropertyIdentifier = 114
	TimeOfStateCountReset                  PropertyIdentifier = 115
	TimeSynchronizationRecipients          PropertyIdentifier = 116
	Units                                  PropertyIdentifier = 117
	UpdateInterval                         PropertyIdentifier = 118
	UtcOffset                              PropertyIdentifier = 119
	VendorIdentifier                       PropertyIdentifier = 120
	VendorName                             PropertyIdentifier = 121
	VtClassesSupported                     PropertyIdentifier = 122
	WeeklySchedule                         PropertyIdentifier = 123
	AttemptedSamples                       PropertyIdentifier = 124
	AverageValue                           PropertyIdentifier = 125
	BufferSize                             PropertyIdentifier = 126
	ClientCovIncrement                     PropertyIdentifier = 127
	CovResubscriptionInterval              PropertyIdentifier = 128
	CurrentNotifyTime                      PropertyIdentifier = 129
	EventTimeStamps                        PropertyIdentifier = 130
	LogBuffer                              PropertyIdentifier = 131
	LogDeviceObjectProperty                PropertyIdentifier = 132
	Enable                                 PropertyIdentifier = 133
	LogInterval                            PropertyIdentifier = 134
	MaximumValue                           PropertyIdentifier = 135
	MinimumValue                           PropertyIdentifier = 136
	NotificationThreshold                  PropertyIdentifier = 137
	PreviousNotifyTime                     PropertyIdentifier = 138
	ProtocolRevision                       PropertyIdentifier = 139
	RecordsSinceNotification               PropertyIdentifier = 140
	RecordCount                            PropertyIdentifier = 141
	StartTime                              PropertyIdentifier = 142
	StopTime                               PropertyIdentifier = 143
	StopWhenFull                           PropertyIdentifier = 144
	TotalRecordCount                       PropertyIdentifier = 145
	ValidSamples                           PropertyIdentifier = 146
	WindowInterval                         PropertyIdentifier = 147
	WindowSamples                          PropertyIdentifier = 148
	MaximumValueTimestamp                  PropertyIdentifier = 149
	MinimumValueTimestamp                  PropertyIdentifier = 150
	VarianceValue                          PropertyIdentifier = 151
	ActiveCovSubscriptions                 PropertyIdentifier = 152
	BackupFailureTimeout                   PropertyIdentifier = 153
	ConfigurationFiles                     PropertyIdentifier = 154
	DatabaseRevision                       PropertyIdentifier = 155
	DirectReading                          PropertyIdentifier = 156
	LastRestoreTime                        PropertyIdentifier = 157
	MaintenanceRequired                    PropertyIdentifier = 158
	MemberOf                               PropertyIdentifier = 159
	Mode                                   PropertyIdentifier = 160
	OperationExpected                      PropertyIdentifier = 161
	Setting                                PropertyIdentifier = 162
	Silenced                               PropertyIdentifier = 163
	TrackingValue                          PropertyIdentifier = 164
	ZoneMembers                            PropertyIdentifier = 165
	LifeSafetyAlarmValues                  PropertyIdentifier = 166
	MaxSegmentsAccepted                    PropertyIdentifier = 167
	ProfileName                            PropertyIdentifier = 168
	AutoSlaveDiscovery                     PropertyIdentifier = 169
	ManualSlaveAddressBinding              PropertyIdentifier = 170
	SlaveAddressBinding                    PropertyIdentifier = 171
	SlaveProxyEnable                       PropertyIdentifier = 172
	LastNotifyRecord                       PropertyIdentifier = 173
	ScheduleDefault                        PropertyIdentifier = 174
	AcceptedModes                          PropertyIdentifier = 175
	AdjustValue                            PropertyIdentifier = 176
	Count                                  PropertyIdentifier = 177
	CountBeforeChange                      PropertyIdentifier = 178
	CountChangeTime                        PropertyIdentifier = 179
	CovPeriod                              PropertyIdentifier = 180
	InputReference                         PropertyIdentifier = 181
	LimitMonitoringInterval                PropertyIdentifier = 182
	LoggingObject                          PropertyIdentifier = 183
	LoggingRecord                          PropertyIdentifier = 184
	Prescale                               PropertyIdentifier = 185
	PulseRate                              PropertyIdentifier = 186
	Scale                                  PropertyIdentifier = 187
	ScaleFactor                            PropertyIdentifier = 188
	UpdateTime                             PropertyIdentifier = 189
	ValueBeforeChange                      PropertyIdentifier = 190
	ValueSet                               PropertyIdentifier = 191
	ValueChangeTime                        PropertyIdentifier = 192
	AlignIntervals                         PropertyIdentifier = 193
	IntervalOffset                         PropertyIdentifier = 195
	LastRestartReason                      PropertyIdentifier = 196
	PropertyIdentifierLoggingType          PropertyIdentifier = 197
	RestartNotificationRecipients          PropertyIdentifier = 202
	TimeOfDeviceRestart                    PropertyIdentifier = 203
	TimeSynchronizationInterval            PropertyIdentifier = 204
	Trigger                                PropertyIdentifier = 205
	UtcTimeSynchronizationRecipients       PropertyIdentifier = 206
	NodeSubtype                            PropertyIdentifier = 207
	PropertyIdentifierNodeType             PropertyIdentifier = 208
	StructuredObjectList                   PropertyIdentifier = 209
	SubordinateAnnotations                 PropertyIdentifier = 210
	SubordinateList                        PropertyIdentifier = 211
	ActualShedLevel                        PropertyIdentifier = 212
	DutyWindow                             PropertyIdentifier = 213
	ExpectedShedLevel                      PropertyIdentifier = 214
	FullDutyBaseline                       PropertyIdentifier = 215
	RequestedShedLevel                     PropertyIdentifier = 218
	ShedDuration                           PropertyIdentifier = 219
	ShedLevelDescriptions                  PropertyIdentifier = 220
	ShedLevels                             PropertyIdentifier = 221
	StateDescription                       PropertyIdentifier = 222
	PropertyIdentifierDoorAlarmState       PropertyIdentifier = 226
	DoorExtendedPulseTime                  PropertyIdentifier = 227
	DoorMembers                            PropertyIdentifier = 228
	DoorOpenTooLongTime                    PropertyIdentifier = 229
	DoorPulseTime                          PropertyIdentifier = 230
	PropertyIdentifierDoorStatus           PropertyIdentifier = 231
	DoorUnlockDelayTime                    PropertyIdentifier = 232
	PropertyIdentifierLockStatus           PropertyIdentifier = 233
	MaskedAlarmValues                      PropertyIdentifier = 234
	SecuredStatus                          PropertyIdentifier = 235
	AbsenteeLimit                          PropertyIdentifier = 244
	AccessAlarmEvents                      PropertyIdentifier = 245
	AccessDoors                            PropertyIdentifier = 246
	PropertyIdentifierAccessEvent          PropertyIdentifier = 247
	AccessEventAuthenticationFactor        PropertyIdentifier = 248
	AccessEventCredential                  PropertyIdentifier = 249
	AccessEventTime                        PropertyIdentifier = 250
	AccessTransactionEvents                PropertyIdentifier = 251
	Accompaniment                          PropertyIdentifier = 252
	AccompanimentTime                      PropertyIdentifier = 253
	ActivationTime                         PropertyIdentifier = 254
	ActiveAuthenticationPolicy             PropertyIdentifier = 255
	AssignedAccessRights                   PropertyIdentifier = 256
	AuthenticationFactors                  PropertyIdentifier = 257
	AuthenticationPolicyList               PropertyIdentifier = 258
	AuthenticationPolicyNames              PropertyIdentifier = 259
	PropertyIdentifierAuthenticationStatus PropertyIdentifier = 260
	PropertyIdentifierAuthorizationMode    PropertyIdentifier = 261
	BelongsTo                              PropertyIdentifier = 262
	CredentialDisable                      PropertyIdentifier = 263
	CredentialStatus                       PropertyIdentifier = 264
	Credentials                            PropertyIdentifier = 265
	CredentialsInZone                      PropertyIdentifier = 266
	DaysRemaining                          PropertyIdentifier = 267
	EntryPoints                            PropertyIdentifier = 268
	ExitPoints                             PropertyIdentifier = 269
	ExpiryTime                             PropertyIdentifier = 270
	ExtendedTimeEnable                     PropertyIdentifier = 271
	FailedAttemptEvents                    PropertyIdentifier = 272
	FailedAttempts                         PropertyIdentifier = 273
	FailedAttemptsTime                     PropertyIdentifier = 274
	LastAccessEvent                        PropertyIdentifier = 275
	LastAccessPoint                        PropertyIdentifier = 276
	LastCredentialAdded                    PropertyIdentifier = 277
	LastCredentialAddedTime                PropertyIdentifier = 278
	LastCredentialRemoved                  PropertyIdentifier = 279
	LastCredentialRemovedTime              PropertyIdentifier = 280
	LastUseTime                            PropertyIdentifier = 281
	Lockout                                PropertyIdentifier = 282
	LockoutRelinquishTime                  PropertyIdentifier = 283
	MasterExemption                        PropertyIdentifier = 284
	MaxFailedAttempts                      PropertyIdentifier = 285
	Members                                PropertyIdentifier = 286
	MusterPoint                            PropertyIdentifier = 287
	NegativeAccessRules                    PropertyIdentifier = 288
	NumberOfAuthenticationPolicies         PropertyIdentifier = 289
	OccupancyCount                         PropertyIdentifier = 290
	OccupancyCountAdjust                   PropertyIdentifier = 291
	OccupancyCountEnable                   PropertyIdentifier = 292
	OccupancyExemption                     PropertyIdentifier = 293
	OccupancyLowerLimit                    PropertyIdentifier = 294
	OccupancyLowerLimitEnforced            PropertyIdentifier = 295
	OccupancyState                         PropertyIdentifier = 296
	OccupancyUpperLimit                    PropertyIdentifier = 297
	OccupancyUpperLimitEnforced            PropertyIdentifier = 298
	PassbackExemption                      PropertyIdentifier = 299
	PassbackMode                           PropertyIdentifier = 300
	PassbackTimeout                        PropertyIdentifier = 301
	PositiveAccessRules                    PropertyIdentifier = 302
	ReasonForDisable                       PropertyIdentifier = 303
	SupportedFormats                       PropertyIdentifier = 304
	SupportedFormatClasses                 PropertyIdentifier = 305
	ThreatAuthority                        PropertyIdentifier = 306
	ThreatLevel                            PropertyIdentifier = 307
	TraceFlag                              PropertyIdentifier = 308
	TransactionNotificationClass           PropertyIdentifier = 309
	UserExternalIdentifier                 PropertyIdentifier = 310
	UserInformationReference               PropertyIdentifier = 311
	UserName                               PropertyIdentifier = 317
	UserType                               PropertyIdentifier = 318
	UsesRemaining                          PropertyIdentifier = 319
	ZoneFrom                               PropertyIdentifier = 320
	ZoneTo                                 PropertyIdentifier = 321
	AccessEventTag                         PropertyIdentifier = 322
	GlobalIdentifier                       PropertyIdentifier = 323
	VerificationTime                       PropertyIdentifier = 326
	BaseDeviceSecurityPolicy               PropertyIdentifier = 327
	DistributionKeyRevision                PropertyIdentifier = 328
	DoNotHide                              PropertyIdentifier = 329
	KeySets                                PropertyIdentifier = 330
	LastKeyServer                          PropertyIdentifier = 331
	NetworkAccessSecurityPolicies          PropertyIdentifier = 332
	PacketReorderTime                      PropertyIdentifier = 333
	SecurityPduTimeout                     PropertyIdentifier = 334
	SecurityTimeWindow                     PropertyIdentifier = 335
	SupportedSecurityAlgorithm             PropertyIdentifier = 336
	UpdateKeySetTimeout                    PropertyIdentifier = 337
	BackupAndRestoreState                  PropertyIdentifier = 338
	BackupPreparationTime                  PropertyIdentifier = 339
	RestoreCompletionTime                  PropertyIdentifier = 340
	RestorePreparationTime                 PropertyIdentifier = 341
	BitMask                                PropertyIdentifier = 342
	BitText                                PropertyIdentifier = 343
	IsUtc                                  PropertyIdentifier = 344
	GroupMembers                           PropertyIdentifier = 345
	GroupMemberNames                       PropertyIdentifier = 346
	MemberStatusFlags                      PropertyIdentifier = 347
	RequestedUpdateInterval                PropertyIdentifier = 348
	CovuPeriod                             PropertyIdentifier = 349
	CovuRecipients                         PropertyIdentifier = 350
	EventMessageTexts                      PropertyIdentifier = 351
	EventMessageTextsConfig                PropertyIdentifier = 352
	EventDetectionEnable                   PropertyIdentifier = 353
	EventAlgorithmInhibit                  PropertyIdentifier = 354
	EventAlgorithmInhibitRef               PropertyIdentifier = 355
	TimeDelayNormal                        PropertyIdentifier = 356
	ReliabilityEvaluationInhibit           PropertyIdentifier = 357
	FaultParameters                        PropertyIdentifier = 358
	PropertyIdentifierFaultType            PropertyIdentifier = 359
	LocalForwardingOnly                    PropertyIdentifier = 360
	ProcessIdentifierFilter                PropertyIdentifier = 361
	SubscribedRecipients                   PropertyIdentifier = 362
	PortFilter                             PropertyIdentifier = 363
	AuthorizationExemptions                PropertyIdentifier = 364
	AllowGroupDelayInhibit                 PropertyIdentifier = 365
	ChannelNumber                          PropertyIdentifier = 366
	ControlGroups                          PropertyIdentifier = 367
	ExecutionDelay                         PropertyIdentifier = 368
	LastPriority                           PropertyIdentifier = 369
	PropertyIdentifierWriteStatus          PropertyIdentifier = 370
	PropertyList                           PropertyIdentifier = 371
	SerialNumber                           PropertyIdentifier = 372
	BlinkWarnEnable                        PropertyIdentifier = 373
	DefaultFadeTime                        PropertyIdentifier = 374
	DefaultRampRate                        PropertyIdentifier = 375
	DefaultStepIncrement                   PropertyIdentifier = 376
	EgressTime                             PropertyIdentifier = 377
	InProgress                             PropertyIdentifier = 378
	InstantaneousPower                     PropertyIdentifier = 379
	LightingCommand                        PropertyIdentifier = 380
	LightingCommandDefaultPriority         PropertyIdentifier = 381
	MaxActualValue                         PropertyIdentifier = 382
	MinActualValue                         PropertyIdentifier = 383
	Power                                  PropertyIdentifier = 384
	Transition                             PropertyIdentifier = 385
	EgressActive                           PropertyIdentifier = 386
	InterfaceValue                         PropertyIdentifier = 387
	FaultHighLimit                         PropertyIdentifier = 388
	FaultLowLimit                          PropertyIdentifier = 389
	LowDiffLimit                           PropertyIdentifier = 390
	StrikeCount                            PropertyIdentifier = 391
	TimeOfStrikeCountReset                 PropertyIdentifier = 392
	DefaultTimeout                         PropertyIdentifier = 393
	InitialTimeout                         PropertyIdentifier = 394
	LastStateChange                        PropertyIdentifier = 395
	StateChangeValues                      PropertyIdentifier = 396
	TimerRunning                           PropertyIdentifier = 397
	PropertyIdentifierTimerState           PropertyIdentifier = 398
	ApduLength                             PropertyIdentifier = 399
	IpAddress                              PropertyIdentifier = 400
	IpDefaultGateway                       PropertyIdentifier = 401
	IpDhcpEnable                           PropertyIdentifier = 402
	IpDhcpLeaseTime                        PropertyIdentifier = 403
	IpDhcpLeaseTimeRemaining               PropertyIdentifier = 404
	IpDhcpServer                           PropertyIdentifier = 405
	IpDnsServer                            PropertyIdentifier = 406
	BacnetIpGlobalAddress                  PropertyIdentifier = 407
	BacnetIpMode                           PropertyIdentifier = 408
	BacnetIpMulticastAddress               PropertyIdentifier = 409
	BacnetIpNatTraversal                   PropertyIdentifier = 410
	IpSubnetMask                           PropertyIdentifier = 411
	BacnetIpUdpPort                        PropertyIdentifier = 412
	BbmdAcceptFdRegistrations              PropertyIdentifier = 413
	BbmdBroadcastDistributionTable         PropertyIdentifier = 414
	BbmdForeignDeviceTable                 PropertyIdentifier = 415
	ChangesPending                         PropertyIdentifier = 416
	Command                                PropertyIdentifier = 417
	FdBbmdAddress                          PropertyIdentifier = 418
	FdSubscriptionLifetime                 PropertyIdentifier = 419
	LinkSpeed                              PropertyIdentifier = 420
	LinkSpeeds                             PropertyIdentifier = 421
	LinkSpeedAutonegotiate                 PropertyIdentifier = 422
	MacAddress                             PropertyIdentifier = 423
	NetworkInterfaceName                   PropertyIdentifier = 424
	NetworkNumber                          PropertyIdentifier = 425
	PropertyIdentifierNetworkNumberQuality PropertyIdentifier = 426
	PropertyIdentifierNetworkType          PropertyIdentifier = 427
	RoutingTable                           PropertyIdentifier = 428
	VirtualMacAddressTable                 PropertyIdentifier = 429
	CommandTimeArray                       PropertyIdentifier = 430
	CurrentCommandPriority                 PropertyIdentifier = 431
	LastCommandTime                        PropertyIdentifier = 432
	ValueSource                            PropertyIdentifier = 433
	ValueSourceArray                       PropertyIdentifier = 434
	BacnetIpv6Mode                         PropertyIdentifier = 435
	Ipv6Address                            PropertyIdentifier = 436
	Ipv6PrefixLength                       PropertyIdentifier = 437
	BacnetIpv6UdpPort                      PropertyIdentifier = 438
	Ipv6DefaultGateway                     PropertyIdentifier = 439
	BacnetIpv6MulticastAddress             PropertyIdentifier = 440
	Ipv6DnsServer                          PropertyIdentifier = 441
	Ipv6AutoAddressingEnable               PropertyIdentifier = 442
	Ipv6DhcpLeaseTime                      PropertyIdentifier = 443
	Ipv6DhcpLeaseTimeRemaining             PropertyIdentifier = 444
	Ipv6DhcpServer                         PropertyIdentifier = 445
	Ipv6ZoneIndex                          PropertyIdentifier = 446
	AssignedLandingCalls                   PropertyIdentifier = 447
	CarAssignedDirection                   PropertyIdentifier = 448
	CarDoorCommand                         PropertyIdentifier = 449
	CarDoorStatus                          PropertyIdentifier = 450
	CarDoorText                            PropertyIdentifier = 451
	CarDoorZone                            PropertyIdentifier = 452
	CarDriveStatus                         PropertyIdentifier = 453
	CarLoad                                PropertyIdentifier = 454
	CarLoadUnits                           PropertyIdentifier = 455
	CarMode                                PropertyIdentifier = 456
	CarMovingDirection                     PropertyIdentifier = 457
	CarPosition                            PropertyIdentifier = 458
	ElevatorGroup                          PropertyIdentifier = 459
	EnergyMeter                            PropertyIdentifier = 460
	EnergyMeterRef                         PropertyIdentifier = 461
	PropertyIdentifierEscalatorMode        PropertyIdentifier = 462
	FloorText                              PropertyIdentifier = 464
	GroupId                                PropertyIdentifier = 465
	GroupMode                              PropertyIdentifier = 467
	HigherDeck                             PropertyIdentifier = 468
	InstallationId                         PropertyIdentifier = 469
	LandingCalls                           PropertyIdentifier = 470
	LandingCallControl                     PropertyIdentifier = 471
	LandingDoorStatus                      PropertyIdentifier = 472
	LowerDeck                              PropertyIdentifier = 473
	MachineRoomId                          PropertyIdentifier = 474
	MakingCarCall                          PropertyIdentifier = 475
	NextStoppingFloor                      PropertyIdentifier = 476
	OperationDirection                     PropertyIdentifier = 477
	PassengerAlarm                         PropertyIdentifier = 478
	PowerMode                              PropertyIdentifier = 479
	RegisteredCarCall                      PropertyIdentifier = 480
	ActiveCovMultipleSubscriptions         PropertyIdentifier = 481
	PropertyIdentifierProtocolLevel        PropertyIdentifier = 482
	ReferencePort                          PropertyIdentifier = 483
	DeployedProfileLocation                PropertyIdentifier = 484
	ProfileLocation                        PropertyIdentifier = 485
	Tags                                   PropertyIdentifier = 486
	SubordinateNodeTypes                   PropertyIdentifier = 487
	SubordinateRelationships               PropertyIdentifier = 489
	SubordinateTags                        PropertyIdentifier = 488
	DefaultSubordinateRelationship         PropertyIdentifier = 490
	Represents                             PropertyIdentifier = 491
	DefaultPresentValue                    PropertyIdentifier = 492
	PresentStage                           PropertyIdentifier = 493
	Stages                                 PropertyIdentifier = 494
	StageNames                             PropertyIdentifier = 495
	TargetReferences                       PropertyIdentifier = 496
	FaultSignals                           PropertyIdentifier = 463
)

type Segmentation int

const (
	SegmentedBoth Segmentation = iota
	SegmentedTransmit
	SegmentedReceive
	NoSegmentation
)

type EventType int

const (
	ChangeOfBitstring          EventType = 0
	ChangeOfState              EventType = 1
	ChangeOfValue              EventType = 2
	CommandFailure             EventType = 3
	FloatingLimit              EventType = 4
	OutOfRange                 EventType = 5
	Complex                    EventType = 6
	ChangeOfLifeSafety         EventType = 8
	Extended                   EventType = 9
	BufferReady                EventType = 10
	UnsignedRange              EventType = 11
	BACnetEventTypeAccessEvent EventType = 13
	DoubleOutOfRange           EventType = 14
	SignedOutOfRange           EventType = 15
	UnsignedOutOfRange         EventType = 16
	ChangeOfCharacterstring    EventType = 17
	ChangeOfStatusFlag         EventType = 18
	ChangeOfReliability        EventType = 19
	None                       EventType = 20
	ChangeOfDiscreteValue      EventType = 21
	ChangeOfTimer              EventType = 22
)

type FaultType int

const (
	BacnetFaultTypeNone FaultType = iota
	FaultCHARACTERSTRING
	FaultEXTENDED
	FaultLIFE_SAFETY
	FaultSTATE
	FaultStatusFlags
	FaultOutOfRange
	FaultListed
)

type NotifyType int

const (
	ALARM NotifyType = iota
	EVENT
	ACK_NOTIFICATION
)

type EventState int

const (
	Normal EventState = iota
	Fault
	OffNormal
	BACnetEventStateHighLimit
	BACnetEventStateLowLimit
	LifeSafetyAlarm
)

type AccessCredentialDisableReason int

const (
	Disabled AccessCredentialDisableReason = iota
	DisabledNeedsProvisioning
	DisabledUnassigned
	DisabledNotYetActive
	DisabledExpired
	DisabledLockout
	DisabledMaxDays
	DisabledMaxUses
	DisabledInactivity
	DisabledManual
)

type AccessCredentialDisable int

const (
	BACnetAccessCredentialDisableNone AccessCredentialDisable = iota
	Disable
	DisableManual
	DisableLockout
)

type AccessPassbackMode int

const (
	PassbackOff AccessPassbackMode = iota
	HardPassback
	SoftPassback
)

type AccessUserType int

const (
	Asset AccessUserType = iota
	BACnetAccessUserTypeGroup
	Person
)

type AccessZoneOccupancyState int

const (
	BACnetAccessZoneOccupancyStateNormal AccessZoneOccupancyState = iota
	BelowLowerLimit
	AtLowerLimit
	AtUpperLimit
	AboveUpperLimit
	BACnetAccessZoneOccupancyStateDisabled
	NotSupported
)

type Action int

const (
	Direct Action = iota
	Reverse
)

type NetworkNumberQuality int

const (
	Unknown NetworkNumberQuality = iota
	Learned
	LearnedConfigured
	Configured
)

type BinaryPV int

const (
	Inactive BinaryPV = iota
	Active
)

type DoorValue int

const (
	Lock DoorValue = iota
	Unlock
	PulseUnlock
	ExtendedPulseUnlock
)

type AuthenticationStatus int

const (
	NotReady AuthenticationStatus = iota
	Ready
	BACnetAuthenticationStatusDisabled
	WaitingForAuthenticationFactor
	WaitingForAccompaniment
	WaitingForVerification
	BACnetAuthenticationStatusInProgress
)

type AuthorizationExemption int

const (
	Passback AuthorizationExemption = iota
	OccupancyCheck
	BACnetAuthorizationExemptionAccessRights
	BACnetAuthorizationExemptionLockout
	Deny
	Verification
	AuthorizationDelay
)

type AuthorizationMode int

const (
	Authorize AuthorizationMode = iota
	GrantActive
	DenyAll
	VerificationRequired
	AuthorizationDelayed
	BACnetAuthorizationModeNone
)

type BackupState int

const (
	Idle BackupState = iota
	PreparingForBackup
	PreparingForRestor
	PerformingABACKUP
	PerformingARestor
)

type BinaryLightingPV int

const (
	Off BinaryLightingPV = iota
	On
	Warn
	WarnOff
	WarnRelinquish
	Stop
)

type DeviceStatus int

const (
	Operational DeviceStatus = iota
	OperationalReadOnly
	DownloadRequired
	DownloadInProgress
	NonOperational
	BackupInProgress
)

type DoorAlarmState int

const (
	BACnetDoorAlarmStateNormal DoorAlarmState = iota
	Alarm
	DoorOpenTooLong
	ForcedOpen
	Tamper
	DoorFault
	LockDown
	FreeAccess
	EgressOpen
)

type DoorSecuredStatus int

const (
	Secured DoorSecuredStatus = iota
	UNSecured
	BACnetDoorSecuredStatusUnknown
)

type DoorStatus int

const (
	CLOSED DoorStatus = iota
	OPENED
	UNKNOWN
	DOOR_FAULT
	UNUSED
	NONE
	CLOSING
	OPENING
	SAFETY_LOCKED
	LIMITED_OPENED
)

type EngineeringUnits int

const (
	metersPerSecondPerSecond          EngineeringUnits = 166
	SquareMeters                      EngineeringUnits = 0
	SquareCentimeters                 EngineeringUnits = 116
	SquareFeet                        EngineeringUnits = 1
	SquareInches                      EngineeringUnits = 115
	Currency1                         EngineeringUnits = 105
	Currency2                         EngineeringUnits = 106
	Currency3                         EngineeringUnits = 107
	Currency4                         EngineeringUnits = 108
	Currency5                         EngineeringUnits = 109
	Currency6                         EngineeringUnits = 110
	Currency7                         EngineeringUnits = 111
	Currency8                         EngineeringUnits = 112
	Currency9                         EngineeringUnits = 113
	Currency10                        EngineeringUnits = 114
	Milliamperes                      EngineeringUnits = 2
	Amperes                           EngineeringUnits = 3
	AmperesPerMeter                   EngineeringUnits = 167
	AmperesPerSquareMeter             EngineeringUnits = 168
	AmpereSquareMeters                EngineeringUnits = 169
	Decibels                          EngineeringUnits = 199
	DecibelsMillivolt                 EngineeringUnits = 200
	DecibelsVolt                      EngineeringUnits = 201
	Farads                            EngineeringUnits = 170
	Henrys                            EngineeringUnits = 171
	Ohms                              EngineeringUnits = 4
	OhmMeters                         EngineeringUnits = 172
	Milliohms                         EngineeringUnits = 145
	Kilohms                           EngineeringUnits = 122
	Megohms                           EngineeringUnits = 123
	Microsiemens                      EngineeringUnits = 190
	Millisiemens                      EngineeringUnits = 202
	Siemens                           EngineeringUnits = 173
	SiemensPerMeter                   EngineeringUnits = 174
	Teslas                            EngineeringUnits = 175
	Volts                             EngineeringUnits = 5
	Millivolts                        EngineeringUnits = 124
	Kilovolts                         EngineeringUnits = 6
	Megavolts                         EngineeringUnits = 7
	VoltAmperes                       EngineeringUnits = 8
	KilovoltAmperes                   EngineeringUnits = 9
	MegavoltAmperes                   EngineeringUnits = 10
	VoltAmperesReactive               EngineeringUnits = 11
	KilovoltAmperesReactive           EngineeringUnits = 12
	MegavoltAmperesReactive           EngineeringUnits = 13
	VoltsPerDegreeKelvin              EngineeringUnits = 176
	VoltsPerMeter                     EngineeringUnits = 177
	DegreesPhase                      EngineeringUnits = 14
	PowerFactor                       EngineeringUnits = 15
	Webers                            EngineeringUnits = 178
	Joules                            EngineeringUnits = 16
	Kilojoules                        EngineeringUnits = 17
	KilojoulesPerKilogram             EngineeringUnits = 125
	Megajoules                        EngineeringUnits = 126
	WattHours                         EngineeringUnits = 18
	KilowattHours                     EngineeringUnits = 19
	MegawattHours                     EngineeringUnits = 146
	WattHoursReactive                 EngineeringUnits = 203
	KilowattHoursReactive             EngineeringUnits = 204
	MegawattHoursReactive             EngineeringUnits = 205
	Btus                              EngineeringUnits = 20
	KiloBtus                          EngineeringUnits = 147
	MegaBtus                          EngineeringUnits = 148
	Therms                            EngineeringUnits = 21
	TonHours                          EngineeringUnits = 22
	JoulesPerKilogramDryAir           EngineeringUnits = 23
	KilojoulesPerKilogramDryAir       EngineeringUnits = 149
	MegajoulesPerKilogramDryAir       EngineeringUnits = 150
	BtusPerPoundDryAir                EngineeringUnits = 24
	BtusPerPound                      EngineeringUnits = 117
	JoulesPerDegreeKelvin             EngineeringUnits = 127
	KilojoulesPerDegreeKelvin         EngineeringUnits = 151
	MegajoulesPerDegreeKelvin         EngineeringUnits = 152
	JoulesPerKilogramDegreeKelvin     EngineeringUnits = 128
	Newton                            EngineeringUnits = 153
	CyclesPerHour                     EngineeringUnits = 25
	CyclesPerMinute                   EngineeringUnits = 26
	Hertz                             EngineeringUnits = 27
	Kilohertz                         EngineeringUnits = 129
	Megahertz                         EngineeringUnits = 130
	PerHour                           EngineeringUnits = 131
	GramsOfWaterPerKilogramDryAir     EngineeringUnits = 28
	PercentRelativeHumidity           EngineeringUnits = 29
	Micrometers                       EngineeringUnits = 194
	Millimeters                       EngineeringUnits = 30
	Centimeters                       EngineeringUnits = 118
	Kilometers                        EngineeringUnits = 193
	Meters                            EngineeringUnits = 31
	Inches                            EngineeringUnits = 32
	Feet                              EngineeringUnits = 33
	Candelas                          EngineeringUnits = 179
	CandelasPerSquareMeter            EngineeringUnits = 180
	WattsPerSquareFoot                EngineeringUnits = 34
	WattsPerSquareMeter               EngineeringUnits = 35
	Lumens                            EngineeringUnits = 36
	Luxes                             EngineeringUnits = 37
	FootCandles                       EngineeringUnits = 38
	Milligrams                        EngineeringUnits = 196
	Grams                             EngineeringUnits = 195
	Kilograms                         EngineeringUnits = 39
	PoundsMass                        EngineeringUnits = 40
	Tons                              EngineeringUnits = 41
	GramsPerSecond                    EngineeringUnits = 154
	GramsPerMinute                    EngineeringUnits = 155
	KilogramsPerSecond                EngineeringUnits = 42
	KilogramsPerMinute                EngineeringUnits = 43
	KilogramsPerHour                  EngineeringUnits = 44
	PoundsMassPerSecond               EngineeringUnits = 119
	PoundsMassPerMinute               EngineeringUnits = 45
	PoundsMassPerHour                 EngineeringUnits = 46
	TonsPerHour                       EngineeringUnits = 156
	Milliwatts                        EngineeringUnits = 132
	Watts                             EngineeringUnits = 47
	Kilowatts                         EngineeringUnits = 48
	Megawatts                         EngineeringUnits = 49
	BtusPerHour                       EngineeringUnits = 50
	KiloBtusPerHour                   EngineeringUnits = 157
	Horsepower                        EngineeringUnits = 51
	TonsRefrigeration                 EngineeringUnits = 52
	Pascals                           EngineeringUnits = 53
	Hectopascals                      EngineeringUnits = 133
	Kilopascals                       EngineeringUnits = 54
	Millibars                         EngineeringUnits = 134
	Bars                              EngineeringUnits = 55
	PoundsForcePerSquareInch          EngineeringUnits = 56
	MillimetersOfWater                EngineeringUnits = 206
	CentimetersOfWater                EngineeringUnits = 57
	InchesOfWater                     EngineeringUnits = 58
	MillimetersOfMercury              EngineeringUnits = 59
	CentimetersOfMercury              EngineeringUnits = 60
	InchesOfMercury                   EngineeringUnits = 61
	DegreesCelsius                    EngineeringUnits = 62
	DegreesKelvin                     EngineeringUnits = 63
	DegreesKelvinPerHour              EngineeringUnits = 181
	DegreesKelvinPerMinute            EngineeringUnits = 182
	DegreesFahrenheit                 EngineeringUnits = 64
	DegreeDaysCelsius                 EngineeringUnits = 65
	DegreeDaysFahrenheit              EngineeringUnits = 66
	DeltaDegreesFahrenheit            EngineeringUnits = 120
	DeltaDegreesKelvin                EngineeringUnits = 121
	Years                             EngineeringUnits = 67
	Months                            EngineeringUnits = 68
	Weeks                             EngineeringUnits = 69
	Days                              EngineeringUnits = 70
	Hours                             EngineeringUnits = 71
	Minutes                           EngineeringUnits = 72
	Seconds                           EngineeringUnits = 73
	HundredthsSeconds                 EngineeringUnits = 158
	Milliseconds                      EngineeringUnits = 159
	NewtonMeters                      EngineeringUnits = 160
	MillimetersPerSecond              EngineeringUnits = 161
	MillimetersPerMinute              EngineeringUnits = 162
	MetersPerSecond                   EngineeringUnits = 74
	MetersPerMinute                   EngineeringUnits = 163
	MetersPerHour                     EngineeringUnits = 164
	KilometersPerHour                 EngineeringUnits = 75
	FeetPerSecond                     EngineeringUnits = 76
	FeetPerMinute                     EngineeringUnits = 77
	MilesPerHour                      EngineeringUnits = 78
	CubicFeet                         EngineeringUnits = 79
	CubicMeters                       EngineeringUnits = 80
	ImperialGallons                   EngineeringUnits = 81
	Milliliters                       EngineeringUnits = 197
	Liters                            EngineeringUnits = 82
	UsGallons                         EngineeringUnits = 83
	CubicFeetPerSecond                EngineeringUnits = 142
	CubicFeetPerMinute                EngineeringUnits = 84
	MillionCubicFeetPerMinute         EngineeringUnits = 254
	CubicFeetPerHour                  EngineeringUnits = 191
	StandardCubicFeetPerDay           EngineeringUnits = 47808
	MillionStandardCubicFeetPerDay    EngineeringUnits = 47809
	ThousandCubicFeetPerDay           EngineeringUnits = 47810
	ThousandStandardCubicFeetPerDay   EngineeringUnits = 47811
	PoundsMassPerDay                  EngineeringUnits = 47812
	CubicMetersPerSecond              EngineeringUnits = 85
	CubicMetersPerMinute              EngineeringUnits = 165
	CubicMetersPerHour                EngineeringUnits = 135
	ImperialGallonsPerMinute          EngineeringUnits = 86
	MillilitersPerSecond              EngineeringUnits = 198
	LitersPerSecond                   EngineeringUnits = 87
	LitersPerMinute                   EngineeringUnits = 88
	LitersPerHour                     EngineeringUnits = 136
	UsGallonsPerMinute                EngineeringUnits = 89
	UsGallonsPerHour                  EngineeringUnits = 192
	DegreesAngular                    EngineeringUnits = 90
	DegreesCelsiusPerHour             EngineeringUnits = 91
	DegreesCelsiusPerMinute           EngineeringUnits = 92
	DegreesFahrenheitPerHour          EngineeringUnits = 93
	DegreesFahrenheitPerMinute        EngineeringUnits = 94
	JouleSeconds                      EngineeringUnits = 183
	KilogramsPerCubicMeter            EngineeringUnits = 186
	KwHoursPerSquareMeter             EngineeringUnits = 137
	KwHoursPerSquareFoot              EngineeringUnits = 138
	MegajoulesPerSquareMeter          EngineeringUnits = 139
	MegajoulesPerSquareFoot           EngineeringUnits = 140
	NoUnits                           EngineeringUnits = 95
	NewtonSeconds                     EngineeringUnits = 187
	NewtonsPerMeter                   EngineeringUnits = 188
	PartsPerMillion                   EngineeringUnits = 96
	PartsPerBillion                   EngineeringUnits = 97
	Percent                           EngineeringUnits = 98
	PercentObscurationPerFoot         EngineeringUnits = 143
	PercentObscurationPerMeter        EngineeringUnits = 144
	PercentPerSecond                  EngineeringUnits = 99
	PerMinute                         EngineeringUnits = 100
	PerSecond                         EngineeringUnits = 101
	PsiPerDegreeFahrenheit            EngineeringUnits = 102
	Radians                           EngineeringUnits = 103
	RadiansPerSecond                  EngineeringUnits = 184
	RevolutionsPerMinute              EngineeringUnits = 104
	SquareMetersPerNewton             EngineeringUnits = 185
	WattsPerMeterPerDegreeKelvin      EngineeringUnits = 189
	WattsPerSquareMeterDegreeKelvin   EngineeringUnits = 141
	PerMille                          EngineeringUnits = 207
	GramsPerGram                      EngineeringUnits = 208
	KilogramsPerKilogram              EngineeringUnits = 209
	GramsPerKilogram                  EngineeringUnits = 210
	MilligramsPerGram                 EngineeringUnits = 211
	MilligramsPerKilogram             EngineeringUnits = 212
	GramsPerMilliliter                EngineeringUnits = 213
	GramsPerLiter                     EngineeringUnits = 214
	MilligramsPerLiter                EngineeringUnits = 215
	MicrogramsPerLiter                EngineeringUnits = 216
	GramsPerCubicMeter                EngineeringUnits = 217
	MilligramsPerCubicMeter           EngineeringUnits = 218
	MicrogramsPerCubicMeter           EngineeringUnits = 219
	NanogramsPerCubicMeter            EngineeringUnits = 220
	GramsPerCubicCentimeter           EngineeringUnits = 221
	Becquerels                        EngineeringUnits = 222
	Kilobecquerels                    EngineeringUnits = 223
	Megabecquerels                    EngineeringUnits = 224
	Gray                              EngineeringUnits = 225
	Milligray                         EngineeringUnits = 226
	Microgray                         EngineeringUnits = 227
	Sieverts                          EngineeringUnits = 228
	Millisieverts                     EngineeringUnits = 229
	Microsieverts                     EngineeringUnits = 230
	MicrosievertsPerHour              EngineeringUnits = 231
	Millirems                         EngineeringUnits = 47814
	MilliremsPerHour                  EngineeringUnits = 47815
	DecibelsA                         EngineeringUnits = 232
	NephelometricTurbidityUnit        EngineeringUnits = 233
	Ph                                EngineeringUnits = 234
	GramsPerSquareMeter               EngineeringUnits = 235
	MinutesPerDegreeKelvin            EngineeringUnits = 236
	MeterSquaredPerMeter              EngineeringUnits = 237
	AmpereSeconds                     EngineeringUnits = 238
	VoltAmpereHours                   EngineeringUnits = 239
	KilovoltAmpereHours               EngineeringUnits = 240
	MegavoltAmpereHours               EngineeringUnits = 241
	VoltAmpereHoursReactive           EngineeringUnits = 242
	KilovoltAmpereHoursReactive       EngineeringUnits = 243
	MegavoltAmpereHoursReactive       EngineeringUnits = 244
	VoltSquareHours                   EngineeringUnits = 245
	AmpereSquareHours                 EngineeringUnits = 246
	JoulePerHours                     EngineeringUnits = 247
	CubicFeetPerDay                   EngineeringUnits = 248
	CubicMetersPerDay                 EngineeringUnits = 249
	WattHoursPerCubicMeter            EngineeringUnits = 250
	JoulesPerCubicMeter               EngineeringUnits = 251
	MolePercent                       EngineeringUnits = 252
	PascalSeconds                     EngineeringUnits = 253
	MillionStandardCubicFeetPerMinute EngineeringUnits = 254
)

type EscalatorMode int

const (
	BacnetescalatorModeUnknown EscalatorMode = iota
	BacnetescalatorModeStop
	BACnetEscalatorModeUp
	BACnetEscalatorModeDown
	BACnetEscalatorModeInspection
	BacnetescalatorModeOutOfService
)

type EscalatorOperationDirection int

const (
	BacnetEscalatorOperationDirectionUnknown EscalatorOperationDirection = iota
	BACnetEscalatorOperationDirectionStopped
	UpRatedSpeed
	UpReducedSpeed
	DownRatedSpeed
	DownReducedSpeed
)

type FileAccessMethod int

const (
	RecordAccess FileAccessMethod = iota
	StreamAccess
)

type IPMode int

const (
	BacnetIPModeNormal IPMode = iota
	Foreign
	Bbmd
)

type LifeSafetyMode int

const (
	BacnetLifeSafetyModeOff LifeSafetyMode = iota
	Lon
	Test
	Manned
	Unmanned
	Armed
	Disarmed
	Prearmed
	Slow
	Fast
	Disconnected
	Enabled
	BacnetLifeSafetyModeDisabled
	AutomaticReleaseDisabled
	Default
)

type LifeSafetyOperation int

const (
	BacnetLifeSafetyOperationNone LifeSafetyOperation = iota
	Silence
	SilenceAudible
	SilenceVisual
	Reset
	ResetAlarm
	ResetFault
	Unsilence
	UnsilenceAudible
	UnsilenceVisual
)

type LifeSafetyState int

const (
	Quiet LifeSafetyState = iota
	PreAlarm
	BacnetlifesafetystateAlarm
	BacnetlifesafetystateFault
	FaultPreAlarm
	FaultAlarm
	BacnetlifesafetystateNotReady
	BacnetlifesafetystateActive
	BacnetlifesafetystateTamper
	TestAlarm
	TestActive
	TestFault
	TestFaultAlarm
	Holdup
	Duress
	TamperAlarm
	Abnormal
	BacnetlifesafetystateEmergencyPower
	Delayed
	Blocked
	LocalAlarm
	GeneralAlarm
	Supervisory
	TestSupervisory
)

type LiftCarDirection int

const (
	BacnetliftcardirectionUnknown LiftCarDirection = iota
	BacnetliftcardirectionNone
	Stopped
	Up
	Down
	UpAndDown
)

type LiftCarDoorCommand int

const (
	BACnetLiftCarDoorCommandNone LiftCarDoorCommand = iota
	Open
	Close
)

type LiftCarDriveStatus int

const (
	BACnetLiftCarDriveStatusUnknown LiftCarDriveStatus = iota
	Stationary
	Braking
	Accelerate
	Decelerate
	RatedSpeed
	SingleFloorJump
	TwoFloorJump
	ThreeFloorJump
	MultiFloorJump
)

type LiftCarMode int

const (
	BACnetLiftCarModeUnknown LiftCarMode = iota
	BACnetLiftCarModeNormal
	Vip
	Homing
	Parking
	AttendantControl
	FirefighterControl
	BACnetLiftCarModeEmergencyPower
	Inspection
	CabinetRecall
	EarthquakeOperation
	FireOperation
	BACnetLiftCarModeOutOfService
	OccupantEvacuation
)

type LiftFault int

const (
	ControllerFault LiftFault = iota
	DriveAndMotorFault
	GovernorAndSafetyGearFault
	LiftShaftDeviceFault
	PowerSupplyFault
	SafetyInterlockFault
	DoorClosingFault
	DoorOpeningFault
	CarStoppedOutsideLandingZone
	CallButtonStuck
	StartFailure
	ControllerSupplyFault
	SelfTestFailure
	RuntimeLimitExceeded
	PositionLost
	DriveTemperatureExceeded
	LoadMeasurementFault
)

type LiftGroupMode int

const (
	BACnetLiftGroupModeUnknown LiftGroupMode = iota
	BACnetLiftGroupModeNormal
	DownPeak
	TwoWay
	FourWay
	EmergencyPower
	UpPeak
)

type LoggingType int

const (
	Polled LoggingType = iota
	Cov
	Triggered
)

type Maintenance int

const (
	BACnetMaintenanceNone Maintenance = iota
	PeriodicTest
	NeedServiceOperational
	NeedServiceInoperative
)

type NetworkPortCommand int

const (
	BACnetNetworkPortCommandIdle NetworkPortCommand = iota
	DiscardChanges
	RenewFdRegistration
	RestartSlaveDiscovery
	RenewDhcp
	RestartAutonegotiation
	Disconnect
	RestartPort
)

type NodeType int

const (
	BACnetNodeTypeUnknown NodeType = iota
	System
	Network
	BACnetNodeTypeDevice
	Organizational
	Area
	Equipment
	Point
	Collection
	BACnetNodeTypeProperty
	Functional
	BACnetNodeTypeOther
	Subsystem
	Building
	Floor
	Section
	Module
	Tree
	Member
	Protocol
	Room
	Zone
)

type Relationship int

const (
	BACnetRelationshipUnknown Relationship = iota
	BACnetRelationshipDefault
	Contains
	ContainedBy
	Uses
	UsedBy
	Commands
	CommandedBy
	Adjusts
	AdjustedBy
	Ingress
	Egress
	SuppliesAir
	ReceivesAir
	SuppliesHotAir
	ReceivesHotAir
	SuppliesCoolAir
	ReceivesCoolAir
	SuppliesPower
	ReceivesPower
	SuppliesGas
	ReceivesGas
	SuppliesWater
	ReceivesWater
	SuppliesHotWater
	ReceivesHotWater
	SuppliesCoolWater
	ReceivesCoolWater
	SuppliesSteam
	ReceivesSteam
)

type Reliability int

const (
	NoFaultDetected Reliability = iota
	NoSensor
	OverRange
	UnderRange
	OpenLoop
	ShortedLoop
	NoOutput
	UnreliableOther
	ProcessError
	MultiStateFault
	ConfigurationError
	CommunicationFailure Reliability = iota + 1
	MemberFault
	MonitoredObjectFault
	Tripped
	LampFailure
	ActivationFailure
	RenewDhcpFailure
	RenewFdRegistrationFailure
	RestartAutoNegotiationFailure
	RestartFailure
	ProprietaryCommandFailure
	FaultsListed
	ReferencedObjectFault
)

type RestartReason int

const (
	BACnetRestartReasonUnknown RestartReason = iota
	ColdStart
	WarmStart
	DetectedPowerLost
	DetectedPowerOff
	HardwareWatchdog
	SoftwareWatchdog
	Suspended
)

type SecurityLevel int

const (
	Incapable SecurityLevel = iota
	Plain
	Signed
	Encrypted
	SignedEndToEnd
	EncryptedEndToEnd
)

type Polarity int

const (
	BACnetPolarityNormal Polarity = iota
	BACnetPolarityReverse
)

type ProtocolLevel int

const (
	Physical ProtocolLevel = iota
	BACnetProtocolLevelProtocol
	BACnetApplication
	NonBACnetApplication
)

type SilencedState int

const (
	Unsilenced SilencedState = iota
	AudibleSilenced
	VisibleSilenced
	AllSilenced
)

type TimerState int

const (
	BACnetTimerStateIdle TimerState = iota
	Running
	Expired
)

type TimerTransition int

const (
	BACnetTimerTransitionNone TimerTransition = iota
	IdleToRunning
	RunningToIdle
	RunningToRunning
	RunningToExpired
	ForcedToExpired
	ExpiredToIdle
	ExpiredToRunning
)

type VTClass int

const (
	DefaultTerminal VTClass = iota
	ANSI_X3_64
	DEC_VT52
	DEC_VT100
	DEC_VT220
	HP_700_94
	IBM_3130
)

type AccessEvent int

const (
	BACnetAccessEventNone AccessEvent = iota
	Granted
	Muster
	PassbackDetected
	BACnetAccessEventDuress
	Trace
	LockoutMaxAttempts
	LockoutOther
	LockoutRelinquished
	LockedByHigherPriority
	BACnetAccessEventOutOfService
	OutOfServiceRelinquished
	AccompanimentBy
	AuthenticationFactorRead
	BACnetAccessEventAuthorizationDelayed
	BACnetAccessEventVerificationRequired
	NoEntryAfterGrant
	DeniedDenyAll AccessEvent = iota + 111
	DeniedUnknownCredential
	DeniedAuthenticationUnavailable
	DeniedAuthenticationFactorTimeout
	DeniedIncorrectAuthenticationFactor
	DeniedZoneNoAccessRights
	DeniedPointNoAccessRights
	DeniedNoAccessRights
	DeniedOutOfTimeRange
	DeniedThreatLevel
	DeniedPassback
	DeniedUnexpectedLocationUsage
	DeniedMaxAttempts
	DeniedLowerOccupancyLimit
	DeniedUpperOccupancyLimit
	DeniedAuthenticationFactorLost
	DeniedAuthenticationFactorStolen
	DeniedAuthenticationFactorDamaged
	DeniedAuthenticationFactorDestroyed
	DeniedAuthenticationFactorDisabled
	DeniedAuthenticationFactorError
	DeniedCredentialUnassigned
	DeniedCredentialNotProvisioned
	DeniedCredentialNotYetActive
	DeniedCredentialExpired
	DeniedCredentialManualDisable
	DeniedCredentialLockout
	DeniedCredentialMaxDays
	DeniedCredentialMaxUses
	DeniedCredentialInactivity
	DeniedCredentialDisabled
	DeniedNoAccompaniment
	DeniedIncorrectAccompaniment
	DeniedLockout
	DeniedVerificationFailed
	DeniedVerificationTimeout
	DeniedOther
)

type LightingInProgress int

const (
	BACnetLightingInProgressIdle LightingInProgress = iota
	FadeActive
	RampActive
	NotControlled
	BACnetLightingInProgressOther
)

type LightingOperation int

const (
	BACnetLightingOperationNone LightingOperation = iota
	FadeTo
	RampTo
	StepUp
	StepDown
	StepOn
	StepOff
	BACnetLightingOperationWarn
	BACnetLightingOperationWarnOff
	BACnetLightingOperationWarnRelinquish
	BACnetLightingOperationStop
)

type LightingTransition int

const (
	BACnetLightingTransitionNone LightingTransition = iota
	Fade
	Ramp
)

type LockStatus int

const (
	Locked LockStatus = iota
	Unlocked
	LockFault
	Unused
	BACnetLockStatusUnknown
)

type EscalatorFault int

const (
	BACnetescalatorfaultControllerFault EscalatorFault = iota
	BACnetescalatorfaultDriveAndMotorFault
	MechanicalComponentFault
	OverspeedFault
	BACnetescalatorfaultPowerSupplyFault
	SafetyDeviceFault
	BACnetescalatorfaultControllerSupplyFault
	BACnetescalatorfaultDriveTemperatureExceeded
	CombPlateFault
)

type ProgramError int

const (
	BACnetProgramErrorNormal = iota
	LoadFailed
	Internal
	BACnetProgramErrorProgram
	BACnetProgramErrorOther
)

type ProgramRequest int

const (
	BACnetProgramRequestReady = iota
	Load
	Run
	Halt
	Restart
	Unload
)

type ProgramState int

const (
	BACnetProgramStateIdle ProgramState = iota
	Loading
	BACnetProgramStateRunning
	Waiting
	Halted
	Unloading
)

type ShedState int

const (
	BACnetShedStateInactive ShedState = iota
	RequestPending
	Compliant
	NonCompliant
)

type WriteStatus int

const (
	BACnetWriteStatusIdle WriteStatus = iota
	BACnetWriteStatusInProgress
	Successful
	Failed
)

type VendorSpecificValue int

func DecodeEnumerated(buffer []byte, offset int, lenValue uint32, objType *ObjectType, propID *PropertyIdentifier) (length int, val interface{}, err error) {
	leng, value, err := DecodeUnsigned(buffer, offset, int(lenValue))
	if err != nil {
		return length, val, err
	}
	if propID != nil {
		switch *propID {
		case SegmentationSupported:
			val = Segmentation(value)
		case PropertyList:
			val = PropertyIdentifier(value)
		case PropertyIdentifierEventType:
			val = EventType(value)
		case PropertyIdentifierNotifyType:
			val = NotifyType(value)
		case PropertyIdentifierFaultType:
			val = FaultType(value)
		case PropertyIdentifierEventState:
			val = EventState(value)
		case ObjectTypePI:
			val = ObjectType(value)
		case ReasonForDisable:
			val = AccessCredentialDisableReason(value)
		case CredentialDisable:
			val = AccessCredentialDisable(value)
		case PassbackMode:
			val = AccessPassbackMode(value)
		case UserType:
			val = AccessUserType(value)
		case PropertyIdentifierNetworkNumberQuality:
			val = NetworkNumberQuality(value)
		case OccupancyState:
			val = AccessZoneOccupancyState(value)
		case PropertyIdentifierAction:
			if *objType == Loop {
				val = Action(value)
			}
		case PresentValue, AlarmValue, FeedbackValue, RelinquishDefault:
			switch *objType {
			case BinaryInput, BinaryOutput, BinaryValue:
				val = BinaryPV(value)
			case AccessDoor:
				val = DoorValue(value)
			case LifeSafetyPoint, LifeSafetyZone:
				val = LifeSafetyState(value)
			case LightingOutput:
				val = BinaryLightingPV(value)
			case LoadControl:
				val = ShedState(value)
			}
		case PropertyIdentifierAuthenticationStatus:
			val = AuthenticationStatus(value)
		case AuthorizationExemptions:
			val = AuthorizationExemption(value)
		case PropertyIdentifierAuthorizationMode:
			val = AuthorizationMode(value)
		case BackupAndRestoreState:
			val = BackupState(value)
		case SystemStatus:
			val = DeviceStatus(value)
		case SecuredStatus:
			val = DoorSecuredStatus(value)
		case PropertyIdentifierDoorStatus, CarDoorStatus:
			val = DoorStatus(value)
		case Units, CarLoadUnits:
			val = EngineeringUnits(value)
		case PropertyIdentifierEscalatorMode:
			val = EscalatorMode(value)
		case OperationDirection:
			val = EscalatorOperationDirection(value)
		case PropertyIdentifierFileAccessMethod:
			val = FileAccessMethod(value)
		case OperationExpected:
			val = LifeSafetyOperation(value)
		case CarDoorCommand:
			val = LiftCarDoorCommand(value)
		case CarDriveStatus:
			val = LiftCarDriveStatus(value)
		case CarMode:
			val = LiftCarMode(value)
		case GroupMode:
			val = LiftGroupMode(value)
		case PropertyIdentifierLoggingType:
			val = LoggingType(value)
		case PropertyIdentifierReliability:
			val = Reliability(value)
		case LastRestartReason:
			val = RestartReason(value)
		case PropertyIdentifierNetworkType:
			val = NetworkType(value)
		case BaseDeviceSecurityPolicy:
			val = SecurityLevel(value)
		case CarMovingDirection, CarAssignedDirection:
			val = LiftCarDirection(value)
		case BacnetIpMode, BacnetIpv6Mode:
			val = IPMode(value)
		case MaintenanceRequired:
			val = Maintenance(value)
		case PropertyIdentifierPolarity:
			val = Polarity(value)
		case PropertyIdentifierProtocolLevel:
			val = ProtocolLevel(value)
		case Silenced:
			val = SilencedState(value)
		case PropertyIdentifierAccessEvent, AccessAlarmEvents, AccessTransactionEvents, FailedAttemptEvents:
			if *objType == AccessPoint {
				val = AccessEvent(value)
			}
		case LastAccessEvent:
			if *objType == AccessCredential {
				val = AccessEvent(value)
			}
		case CredentialStatus:
			if *objType == AccessCredential {
				val = BinaryPV(value)
			}
		case PropertyIdentifierLockStatus:
			if *objType == AccessDoor {
				val = LockStatus(value)
			}
		case PropertyIdentifierDoorAlarmState, MaskedAlarmValues, AlarmValues, FaultValues:
			switch *objType {
			case AccessDoor:
				val = DoorAlarmState(value)
			case LifeSafetyPoint, LifeSafetyZone:
				val = LifeSafetyState(value)
			case Timer:
				val = TimerState(value)
			}
		case Mode, AcceptedModes:
			if *objType == LifeSafetyPoint || *objType == LifeSafetyZone {
				val = LifeSafetyMode(value)
			}
		case TrackingValue, LifeSafetyAlarmValues:
			if *objType == LifeSafetyPoint || *objType == LifeSafetyZone {
				val = LifeSafetyState(value)
			}
		case FaultSignals:
			switch *objType {
			case Escalator:
				val = EscalatorFault(value)
			case Lift:
				val = LiftFault(value)
			}
		case InProgress:
			if *objType == LightingOutput {
				val = LightingInProgress(value)
			}
		case Transition:
			if *objType == LightingOutput {
				val = LightingTransition(value)
			}
		case Command:
			if *objType == NetworkPort {
				val = NetworkPortCommand(value)
			}
		case PropertyIdentifierNodeType, SubordinateNodeTypes:
			if *objType == StructuredView {
				val = NodeType(value)
			}
		case SubordinateRelationships, DefaultSubordinateRelationship:
			if *objType == StructuredView {
				val = Relationship(value)
			}
		case ReasonForHalt:
			if *objType == Program {
				val = ProgramError(value)
			}
		case ProgramChange:
			if *objType == Program {
				val = ProgramRequest(value)
			}
		case PropertyIdentifierProgramState:
			if *objType == Program {
				val = ProgramState(value)
			}
		case PropertyIdentifierTimerState:
			if *objType == Timer {
				val = TimerState(value)
			}
		case LastStateChange:
			if *objType == Timer {
				val = TimerTransition(value)
			}
		case VtClassesSupported:
			val = VTClass(value)
		case PropertyIdentifierWriteStatus:
			if *objType == Channel {
				val = WriteStatus(value)
			}
		default:
			val = VendorSpecificValue(value)
		}

		return leng, val, nil
	}
	return leng, value, nil
}

func EncodeContextEnumerated(tagNumber BACnetApplicationTag, value uint32) []byte {
	length := 0
	if value < 0x100 {
		length = 1
	} else if value < 0x10000 {
		length = 2
	} else if value < 0x1000000 {
		length = 3
	} else {
		length = 4
	}

	return append(EncodeTag(tagNumber, true, length), EncodeUnsigned(value)...)
}
