// Copyright (c) 2013-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package txscript

import (
	"fmt"
)

// ErrorCode identifies a kind of script error.
type ErrorCode int

// These constants are used to identify a specific Error.
const (
	// ErrInternal is returned if internal consistency checks fail.  In
	// practice this error should never be seen as it would mean there is an
	// error in the engine logic.
	ErrInternal ErrorCode = iota

	// ---------------------------------------
	// Failures related to improper API usage.
	// ---------------------------------------

	// ErrInvalidFlags is returned when the passed flags to NewEngine
	// contain an invalid combination.
	ErrInvalidFlags

	// ErrInvalidIndex is returned when an out-of-bounds index is passed to
	// a function.
	ErrInvalidIndex

	// ErrUnsupportedAddress is returned when a concrete type that
	// implements a bchutil.Address is not a supported type.
	ErrUnsupportedAddress

	// ErrNotMultisigScript is returned from CalcMultiSigStats when the
	// provided script is not a multisig script.
	ErrNotMultisigScript

	// ErrTooManyRequiredSigs is returned from MultiSigScript when the
	// specified number of required signatures is larger than the number of
	// provided public keys.
	ErrTooManyRequiredSigs

	// ErrTooMuchNullData is returned from NullDataScript when the length of
	// the provided data exceeds MaxDataCarrierSize.
	ErrTooMuchNullData

	// ------------------------------------------
	// Failures related to final execution state.
	// ------------------------------------------

	// ErrEarlyReturn is returned when OP_RETURN is executed in the script.
	ErrEarlyReturn

	// ErrEmptyStack is returned when the script evaluated without error,
	// but terminated with an empty top stack element.
	ErrEmptyStack

	// ErrEvalFalse is returned when the script evaluated without error but
	// terminated with a false top stack element.
	ErrEvalFalse

	// ErrScriptUnfinished is returned when CheckErrorCondition is called on
	// a script that has not finished executing.
	ErrScriptUnfinished

	// ErrScriptDone is returned when an attempt to execute an opcode is
	// made once all of them have already been executed.  This can happen
	// due to things such as a second call to Execute or calling Step after
	// all opcodes have already been executed.
	ErrInvalidProgramCounter

	// -----------------------------------------------------
	// Failures related to exceeding maximum allowed limits.
	// -----------------------------------------------------

	// ErrScriptTooBig is returned if a script is larger than MaxScriptSize.
	ErrScriptTooBig

	// ErrElementTooBig is returned if the size of an element to be pushed
	// to the stack is over MaxScriptElementSize.
	ErrElementTooBig

	// ErrTooManyOperations is returned if a script has more than
	// MaxOpsPerScript opcodes that do not push data.
	ErrTooManyOperations

	// ErrStackOverflow is returned when stack and altstack combined depth
	// is over the limit.
	ErrStackOverflow

	// ErrInvalidPubKeyCount is returned when the number of public keys
	// specified for a multsig is either negative or greater than
	// MaxPubKeysPerMultiSig.
	ErrInvalidPubKeyCount

	// ErrInvalidSignatureCount is returned when the number of signatures
	// specified for a multisig is either negative or greater than the
	// number of public keys.
	ErrInvalidSignatureCount

	// ErrNumberTooBig is returned when the argument for an opcode that
	// expects numeric input is larger than the expected maximum number of
	// bytes.  For the most part, opcodes that deal with stack manipulation
	// via offsets, arithmetic, numeric comparison, and boolean logic are
	// those that this applies to.  However, any opcode that expects numeric
	// input may fail with this code.
	ErrNumberTooBig

	// ErrNumberTooSmall is returned when the argument for an opcode that
	// expects numeric input is smaller than the expected maximum number of
	// bytes.  For the most part, opcodes that deal with stack manipulation
	// via offsets, arithmetic, numeric comparison, and boolean logic are
	// those that this applies to.  However, any opcode that expects numeric
	// input may fail with this code.
	ErrNumberTooSmall

	// ErrOpCost is returned when the VM cost limit exceeds.
	ErrOpCost

	// ErrTooManyHashIters is returned when hash iteration limit exceeds
	ErrTooManyHashIters

	// ErrConditionalStackDepth is returned when size of the conditional stack
	// is over MaxConditionalStackDepth
	ErrConditionalStackDepth

	// --------------------------------------------
	// Failures related to verification operations.
	// --------------------------------------------

	// ErrVerify is returned when OP_VERIFY is encountered in a script and
	// the top item on the data stack does not evaluate to true.
	ErrVerify

	// ErrEqualVerify is returned when OP_EQUALVERIFY is encountered in a
	// script and the top item on the data stack does not evaluate to true.
	ErrEqualVerify

	// ErrNumEqualVerify is returned when OP_NUMEQUALVERIFY is encountered
	// in a script and the top item on the data stack does not evaluate to
	// true.
	ErrNumEqualVerify

	// ErrCheckSigVerify is returned when OP_CHECKSIGVERIFY is encountered
	// in a script and the top item on the data stack does not evaluate to
	// true.
	ErrCheckSigVerify

	// ErrCheckSigVerify is returned when OP_CHECKMULTISIGVERIFY is
	// encountered in a script and the top item on the data stack does not
	// evaluate to true.
	ErrCheckMultiSigVerify

	// ErrCheckDataSigVerify is returned when OP_CHECKDATASIGVERIFY is
	// encountered in a script and the top item on the data stack does not
	// evaluate to true.
	ErrCheckDataSigVerify

	// --------------------------------------------
	// Failures related to improper use of opcodes.
	// --------------------------------------------

	// ErrDisabledOpcode is returned when a disabled opcode is encountered
	// in a script.
	ErrDisabledOpcode

	// ErrReservedOpcode is returned when an opcode marked as reserved
	// is encountered in a script.
	ErrReservedOpcode

	// ErrMalformedPush is returned when a data push opcode tries to push
	// more bytes than are left in the script.
	ErrMalformedPush

	// ErrInvalidStackOperation is returned when a stack operation is
	// attempted with a number that is invalid for the current stack size.
	ErrInvalidStackOperation

	// ErrUnbalancedConditional is returned when an OP_ELSE or OP_ENDIF is
	// encountered in a script without first having an OP_IF or OP_NOTIF or
	// the end of script is reached without encountering an OP_ENDIF when
	// an OP_IF or OP_NOTIF was previously encountered.
	ErrUnbalancedConditional

	// ErrInvalidInputLength is returned when an input to an opcode is not
	// the correct length as required by that opcode.
	ErrInvalidInputLength

	// ---------------------------------
	// Failures related to malleability.
	// ---------------------------------

	// ErrMinimalData is returned when the ScriptVerifyMinimalData flag
	// is set and the script contains push operations that do not use
	// the minimal opcode required.
	ErrMinimalData

	// ErrMinimalIf is returned when the ScriptVerifyMinimalIf flag
	// is set and the script contains an if or notif that does not use
	// the minimal if encoding.
	ErrMinimalIf

	// ErrInvalidSigHashType is returned when a signature hash type is not
	// one of the supported types.
	ErrInvalidSigHashType

	// ErrSigTooShort is returned when a signature that should be a
	// canonically-encoded DER signature is too short.
	ErrSigTooShort

	// ErrSigTooLong is returned when a signature that should be a
	// canonically-encoded DER signature is too long.
	ErrSigTooLong

	// ErrSigInvalidSeqID is returned when a signature that should be a
	// canonically-encoded DER signature does not have the expected ASN.1
	// sequence ID.
	ErrSigInvalidSeqID

	// ErrSigInvalidDataLen is returned a signature that should be a
	// canonically-encoded DER signature does not specify the correct number
	// of remaining bytes for the R and S portions.
	ErrSigInvalidDataLen

	// ErrSigMissingSTypeID is returned a signature that should be a
	// canonically-encoded DER signature does not provide the ASN.1 type ID
	// for S.
	ErrSigMissingSTypeID

	// ErrSigMissingSLen is returned when a signature that should be a
	// canonically-encoded DER signature does not provide the length of S.
	ErrSigMissingSLen

	// ErrSigInvalidSLen is returned a signature that should be a
	// canonically-encoded DER signature does not specify the correct number
	// of bytes for the S portion.
	ErrSigInvalidSLen

	// ErrSigInvalidRIntID is returned when a signature that should be a
	// canonically-encoded DER signature does not have the expected ASN.1
	// integer ID for R.
	ErrSigInvalidRIntID

	// ErrSigZeroRLen is returned when a signature that should be a
	// canonically-encoded DER signature has an R length of zero.
	ErrSigZeroRLen

	// ErrSigNegativeR is returned when a signature that should be a
	// canonically-encoded DER signature has a negative value for R.
	ErrSigNegativeR

	// ErrSigTooMuchRPadding is returned when a signature that should be a
	// canonically-encoded DER signature has too much padding for R.
	ErrSigTooMuchRPadding

	// ErrSigInvalidSIntID is returned when a signature that should be a
	// canonically-encoded DER signature does not have the expected ASN.1
	// integer ID for S.
	ErrSigInvalidSIntID

	// ErrSigZeroSLen is returned when a signature that should be a
	// canonically-encoded DER signature has an S length of zero.
	ErrSigZeroSLen

	// ErrSigNegativeS is returned when a signature that should be a
	// canonically-encoded DER signature has a negative value for S.
	ErrSigNegativeS

	// ErrSigTooMuchSPadding is returned when a signature that should be a
	// canonically-encoded DER signature has too much padding for S.
	ErrSigTooMuchSPadding

	// ErrSigHighS is returned when the ScriptVerifyLowS flag is set and the
	// script contains any signatures whose S values are higher than the
	// half order.
	ErrSigHighS

	// ErrNotPushOnly is returned when a script that is required to only
	// push data to the stack performs other operations.  A couple of cases
	// where this applies is for a pay-to-script-hash signature script when
	// bip16 is active and when the ScriptVerifySigPushOnly flag is set.
	ErrNotPushOnly

	// ErrSigNullDummy is returned when the ScriptStrictMultiSig flag is set
	// and a multisig script has anything other than 0 for the extra dummy
	// argument.
	ErrSigNullDummy

	// ErrPubKeyType is returned when the ScriptVerifyStrictEncoding
	// flag is set and the script contains invalid public keys.
	ErrPubKeyType

	// ErrCleanStack is returned when the ScriptVerifyCleanStack flag
	// is set, and after evalution, the stack does not contain only a
	// single element.
	ErrCleanStack

	// ErrNullFail is returned when the ScriptVerifyNullFail flag is
	// set and signatures are not empty on failed checksig or checkmultisig
	// operations.
	ErrNullFail

	// -------------------------------
	// Failures related to soft forks.
	// -------------------------------

	// ErrDiscourageUpgradableNOPs is returned when the
	// ScriptDiscourageUpgradableNops flag is set and a NOP opcode is
	// encountered in a script.
	ErrDiscourageUpgradableNOPs

	// ErrNegativeLockTime is returned when a script contains an opcode that
	// interprets a negative lock time.
	ErrNegativeLockTime

	// ErrUnsatisfiedLockTime is returned when a script contains an opcode
	// that involves a lock time and the required lock time has not been
	// reached.
	ErrUnsatisfiedLockTime

	// ErrInvalidDummy is returned when the schnorr dummy element is not
	// encoded correctly. For example, if it has more bytes than necessary
	// to encode the selected pubkeys.
	ErrInvalidDummy

	// ErrInvalidBitCount is returned when the schnorr dummy element does
	// not contain the correct number of bits.
	ErrInvalidBitCount

	// ErrInputSigChecks is returned when the script exceeds the maximum
	// allowed signature density.
	ErrInputSigChecks

	// ErrIntegerOverflow is returned when a stack operation overflows an int64.
	ErrIntegerOverflow

	// numErrorCodes is the maximum error code number used in tests.  This
	// entry MUST be the last entry in the enum.
	numErrorCodes
)

// Map of ErrorCode values back to their constant names for pretty printing.
var errorCodeStrings = map[ErrorCode]string{
	ErrInternal:                 "ErrInternal",
	ErrInvalidFlags:             "ErrInvalidFlags",
	ErrInvalidIndex:             "ErrInvalidIndex",
	ErrUnsupportedAddress:       "ErrUnsupportedAddress",
	ErrNotMultisigScript:        "ErrNotMultisigScript",
	ErrTooManyRequiredSigs:      "ErrTooManyRequiredSigs",
	ErrTooMuchNullData:          "ErrTooMuchNullData",
	ErrEarlyReturn:              "ErrEarlyReturn",
	ErrEmptyStack:               "ErrEmptyStack",
	ErrEvalFalse:                "ErrEvalFalse",
	ErrScriptUnfinished:         "ErrScriptUnfinished",
	ErrInvalidProgramCounter:    "ErrInvalidProgramCounter",
	ErrScriptTooBig:             "ErrScriptTooBig",
	ErrElementTooBig:            "ErrElementTooBig",
	ErrTooManyOperations:        "ErrTooManyOperations",
	ErrStackOverflow:            "ErrStackOverflow",
	ErrInvalidPubKeyCount:       "ErrInvalidPubKeyCount",
	ErrInvalidSignatureCount:    "ErrInvalidSignatureCount",
	ErrNumberTooBig:             "ErrNumberTooBig",
	ErrNumberTooSmall:           "ErrNumberTooSmall",
	ErrOpCost:                   "ErrOpCost",
	ErrTooManyHashIters:         "ErrTooManyHashIters",
	ErrConditionalStackDepth:    "ErrConditionalStackDepth",
	ErrVerify:                   "ErrVerify",
	ErrEqualVerify:              "ErrEqualVerify",
	ErrNumEqualVerify:           "ErrNumEqualVerify",
	ErrCheckSigVerify:           "ErrCheckSigVerify",
	ErrCheckMultiSigVerify:      "ErrCheckMultiSigVerify",
	ErrCheckDataSigVerify:       "ErrCheckDataSigVerify",
	ErrDisabledOpcode:           "ErrDisabledOpcode",
	ErrReservedOpcode:           "ErrReservedOpcode",
	ErrMalformedPush:            "ErrMalformedPush",
	ErrInvalidStackOperation:    "ErrInvalidStackOperation",
	ErrUnbalancedConditional:    "ErrUnbalancedConditional",
	ErrInvalidInputLength:       "ErrInvalidInputLength",
	ErrMinimalData:              "ErrMinimalData",
	ErrMinimalIf:                "ErrMinimalIf",
	ErrInvalidSigHashType:       "ErrInvalidSigHashType",
	ErrSigTooShort:              "ErrSigTooShort",
	ErrSigTooLong:               "ErrSigTooLong",
	ErrSigInvalidSeqID:          "ErrSigInvalidSeqID",
	ErrSigInvalidDataLen:        "ErrSigInvalidDataLen",
	ErrSigMissingSTypeID:        "ErrSigMissingSTypeID",
	ErrSigMissingSLen:           "ErrSigMissingSLen",
	ErrSigInvalidSLen:           "ErrSigInvalidSLen",
	ErrSigInvalidRIntID:         "ErrSigInvalidRIntID",
	ErrSigZeroRLen:              "ErrSigZeroRLen",
	ErrSigNegativeR:             "ErrSigNegativeR",
	ErrSigTooMuchRPadding:       "ErrSigTooMuchRPadding",
	ErrSigInvalidSIntID:         "ErrSigInvalidSIntID",
	ErrSigZeroSLen:              "ErrSigZeroSLen",
	ErrSigNegativeS:             "ErrSigNegativeS",
	ErrSigTooMuchSPadding:       "ErrSigTooMuchSPadding",
	ErrSigHighS:                 "ErrSigHighS",
	ErrNotPushOnly:              "ErrNotPushOnly",
	ErrSigNullDummy:             "ErrSigNullDummy",
	ErrPubKeyType:               "ErrPubKeyType",
	ErrCleanStack:               "ErrCleanStack",
	ErrNullFail:                 "ErrNullFail",
	ErrDiscourageUpgradableNOPs: "ErrDiscourageUpgradableNOPs",
	ErrNegativeLockTime:         "ErrNegativeLockTime",
	ErrUnsatisfiedLockTime:      "ErrUnsatisfiedLockTime",
	ErrInvalidDummy:             "ErrInvalidDummy",
	ErrInvalidBitCount:          "ErrInvalidBitCount",
	ErrInputSigChecks:           "ErrInputSigChecks",
	ErrIntegerOverflow:          "ErrIntegerOverflow",
}

// String returns the ErrorCode as a human-readable name.
func (e ErrorCode) String() string {
	if s := errorCodeStrings[e]; s != "" {
		return s
	}
	return fmt.Sprintf("Unknown ErrorCode (%d)", int(e))
}

// Error identifies a script-related error.  It is used to indicate three
// classes of errors:
//  1. Script execution failures due to violating one of the many requirements
//     imposed by the script engine or evaluating to false
//  2. Improper API usage by callers
//  3. Internal consistency check failures
//
// The caller can use type assertions on the returned errors to access the
// ErrorCode field to ascertain the specific reason for the error.  As an
// additional convenience, the caller may make use of the IsErrorCode function
// to check for a specific error code.
type Error struct {
	ErrorCode   ErrorCode
	Description string
}

// Error satisfies the error interface and prints human-readable errors.
func (e Error) Error() string {
	return e.Description
}

// scriptError creates an Error given a set of arguments.
func scriptError(c ErrorCode, desc string) Error {
	return Error{ErrorCode: c, Description: desc}
}

// IsErrorCode returns whether or not the provided error is a script error with
// the provided error code.
func IsErrorCode(err error, c ErrorCode) bool {
	serr, ok := err.(Error)
	return ok && serr.ErrorCode == c
}
