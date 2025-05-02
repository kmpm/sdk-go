package nats

const (
	// field names
	fieldURL              = "URL"
	fieldConsumerConfig   = "ConsumerConfig"
	fieldSendSubject      = "SendSubject"
	fieldPullConsumerOpts = "PullConsumerOptions"
	fieldPublishOptions   = "PublishOptions"
	fieldFilterSubjects   = "FilterSubjects"

	// error messages
	messageNoConnection                 = "URL or nats connection must be given."
	messageConflictingConnection        = "URL and nats connection were both given."
	messageNoConsumerConfig             = "No consumer config was given."
	messageNoFilterSubjects             = "No filter subjects were given."
	messageMoreThanOneStream            = "More than one stream for given filter subjects."
	messageNoSendSubject                = "Cannot send without a NATS subject defined."
	messageMoreThanOneConsumerConfig    = "More than one consumer config given."
	messageReceiverOptionsWithoutConfig = "Receiver options given without consumer config."
	messageSenderOptionsWithoutSubject  = "Sender options given without send subject."
)
