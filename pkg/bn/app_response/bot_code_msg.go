package appresponse

var SUCESSCODE = "B0000"
var sucessMsg = "Success"

var FAILCODE = "B9999"
var failMsg = "Failed"

var NOTFOUNDTEMPLATECODE = "B9001"
var notFoundTemplateMsg = "Not found template"

var BOTFOUNDEXISTINHISTORYCODE = "B9002"
var botFoundExistInHistoryMsg = "Bot found exist in history"

var BOTNOTFOUNDEXISTINHISTORYCODE = "B9003"
var botNotFoundExistInHistoryMsg = "Bot not found exist in history"

var BOTNOTREGISTEREDCODE = "B9004"
var botNotRegisteredMsg = "Bot not registered"

var EXCHANGENOTERGISTERCODE = "B9005"
var exchangeNotRegisteredMsg = "Exchange not registered"

var BOTNOTOPENEDCODE = "B9006"
var botNotOpenedMsg = "Bot not opened"

var BOTNOTFOUNDCODE = "B9007"
var botNotFoundMsg = "Bot not found"

var BOTFOUNDCODE = "B9008"
var botFoundMsg = "Bot found"

var BOTMAPPING = map[string]string{
	SUCESSCODE:                    sucessMsg,
	FAILCODE:                      failMsg,
	NOTFOUNDTEMPLATECODE:          notFoundTemplateMsg,
	BOTFOUNDEXISTINHISTORYCODE:    botFoundExistInHistoryMsg,
	BOTNOTFOUNDEXISTINHISTORYCODE: botNotFoundExistInHistoryMsg,
	BOTNOTREGISTEREDCODE:          botNotRegisteredMsg,
	EXCHANGENOTERGISTERCODE:       exchangeNotRegisteredMsg,
	BOTNOTOPENEDCODE:              botNotOpenedMsg,
	BOTNOTFOUNDCODE:               botNotFoundMsg,
	BOTFOUNDCODE:                  botFoundMsg,
}
