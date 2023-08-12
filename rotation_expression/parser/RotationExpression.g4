grammar RotationExpression;
options {
	caseInsensitive = true;
}
rotationExpression:
	inlineRotationExpression
	| multilineRotationExpression;
inlineRotationExpression:
	PREAMBLE ':'? WHITESPACE inlineUserList EOF;
inlineUserList: userSpec (',' WHITESPACE userSpec)*;
multilineRotationExpression:
	PREAMBLE ':' WHITESPACE? multilineUserList;
multilineUserList: multilineUserLine+;
multilineUserLine: MULTILINE_SEP userSpec WHITESPACE?;
userSpec: partialUserSpec | unknownUser | completeUserSpec;
username: USERNAME_PART (WHITESPACE USERNAME_PART)*;
unknownUser: UNKNOWN;
partialUserSpec: username;
completeUserSpec: username WHITESPACE '(' USER_ID ')';
MULTILINE_SEP: '\n' ( '-' | '*') WHITESPACE;
WHITESPACE: ' '+;
fragment ROTATES: 'rotates';
fragment AMONG: 'among';
PREAMBLE: ROTATES WHITESPACE AMONG;
UNKNOWN: 'unknown';
USER_ID: [0-9]+;
USERNAME_PART: [a-z0-9._@\-!]+;