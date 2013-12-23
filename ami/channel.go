package ami

import (
	"errors"
	"strconv"
)

//	AbsoluteTimeout	
//		Set absolute timeout.
//		Hangup a channel after a certain time. Acknowledges set time with Timeout Set message.
//
func AbsoluteTimeout(socket *Socket, actionID, channel string, timeout int) (map[string]string, error) {
	// verify channel and action ID
	if len(channel) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: AbsoluteTimeout",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nTimeout: ",
		strconv.Itoa(timeout),
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	Atxfer
//		Attended transfer.
//
func Atxfer(socket *Socket, actionID, channel, exten, context, priority string) (map[string]string, error) {
	// verify channel and action ID
	if len(channel) == 0 || len(actionID) == 0 ||
		len(exten) == 0 || len(context) == 0 || len(priority) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: AbsoluteTimeout",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nExten: ",
		exten,
		"\r\nContext: ",
		context,
		"\r\nPriority: ",
		priority,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	Bridge
//		Bridge two channels already in the PBX.
//	
func Bridge(socket *Socket, actionID, channel1, channel2 string, tone bool) (map[string]string, error) {
	// verify channel and action ID
	if len(channel1) == 0 || len(channel2) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	t := map[bool]string{false: "no", true: "yes"}

	command := []string{
		"Action: Bridge",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel1: ",
		channel1,
		"\r\nChannel2: ",
		channel2,
		"\r\nTone: ",
		t[tone],
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	CoreShowChannels
//		List currently active channels.
//
func CoreShowChannels(socket *Socket, actionID string) ([]map[string]string, error) {
	return getMessageList(socket,
		"CoreShowChannels",
		actionID,
		"CoreShowChannel",
		"CoreShowChannelsComplete")
}

//
//	ExtensionState
//		Check Extension Status.
//
func ExtensionState(socket *Socket, actionID, exten, context string) (map[string]string, error) {
	// verify action ID
	if len(actionID) == 0 || len(exten) == 0 || len(context) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: ExtensionState",
		"\r\nActionID: ",
		actionID,
		"\r\nExten: ",
		exten,
		"\r\nContext: ",
		context,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	Hangup
//		Hangup channel.
//
func Hangup(socket *Socket, actionID, channel, cause string) (map[string]string, error) {
	// verify action ID
	if len(channel) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: Hangup",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nCause: ",
		cause,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	Originate
//		Originate a call.
//		Generates an outgoing call to a Extension/Context/Priority or Application/Data
//
type OriginateData struct {
	exten       string
	context     string
	priority    int
	application string
	data        string
	timeout     int
	callerid    string
	variable    string
	account     string
	async       string
	codecs      string
}

func Originate(socket *Socket, actionID string, originate OriginateData) (map[string]string, error) {
	// verify action ID
	if len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: Originate",
		"\r\nActionID: ",
		actionID,
		"\r\nExten: ",
		originate.exten,
		"\r\nContext: ",
		originate.context,
		"\n\rPriority: ",
		strconv.Itoa(originate.priority),
		"\r\nApplication: ",
		originate.application,
		"\r\nData: ",
		originate.data,
		"\r\nTimeout: ",
		strconv.Itoa(originate.timeout),
		"\r\nCallerID: ",
		originate.callerid,
		"\r\nVariable: ",
		originate.variable,
		"\r\nAccount: ",
		originate.account,
		"\r\nAsync: ",
		originate.async,
		"\r\nCodecs: ",
		originate.codecs,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	Park
//		Park a channel.
//
func Park(socket *Socket, actionID, channel1, channel2 string, timeout int, parkinglot string) (map[string]string, error) {
	// verify action ID
	if len(channel1) == 0 || len(channel2) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: Park",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel1,
		"\r\nChannel2: ",
		channel2,
		"\r\nTimeout: ",
		strconv.Itoa(timeout),
		"\r\nParkinglot: ",
		parkinglot,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	ParkedCalls
//		List parked calls.
//
func ParkedCalls(socket *Socket, actionID string) ([]map[string]string, error) {
	return getMessageList(socket,
		"ParkedCalls",
		actionID,
		"ParkedCall",
		"ParkedCallsComplete")
}

//
//	PlayDTMF
//		Play DTMF signal on a specific channel.
//
func PlayDTMF(socket *Socket, actionID, channel, digit string) (map[string]string, error) {
	// verify action ID
	if len(channel) == 0 || len(digit) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: PlayDTMF",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nDigit: ",
		digit,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	Redirect
//		Redirect (transfer) a call.
//
func Redirect(socket *Socket, actionID, channel, exten, context, priority string) (map[string]string, error) {
	// verify action ID
	if len(channel) == 0 || len(exten) == 0 ||
		len(context) == 0 || len(priority) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: Redirect",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nExten: ",
		exten,
		"\r\nContext: ",
		context,
		"\r\nPriority: ",
		priority,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	SendText
//		Send text message to channel.
//
func SendText(socket *Socket, actionID, channel, msg string) (map[string]string, error) {
	// verify action ID
	if len(channel) == 0 || len(msg) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: SendText",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nMessage: ",
		msg,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)

}

//
//	Setvar
//		Set a channel variable.
//		Set a global or local channel variable.
//		Note:	If a channel name is not provided then the variable is global.
//
func Setvar(socket *Socket, actionID, channel, variable, value string) (map[string]string, error) {
	// verify action ID
	if len(channel) == 0 || len(variable) == 0 || len(value) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: Setvar",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nVariable: ",
		variable,
		"\r\nValue: ",
		value,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)

}

//
//	Status
//		List channel status.
//		Will return the status information of each channel along with the value for the specified channel variables.
//
func Status(socket *Socket, actionID, channel, variables string) (map[string]string, error) {
	// verify action ID
	if len(channel) == 0 || len(variables) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	command := []string{
		"Action: Status",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nVariables: ",
		variables,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}
