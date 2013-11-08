### Credit Card

* http://en.wikipedia.org/wiki/ISO_8583
* http://jpos.org/blog/2013/11/w55y/ 

> In ISO-8582, when you send an authorization request or a financial request and you don’t get a response, you queue a reversal in a persistent store and forward (SAF) queue. So the next time you contact the server, before sending a new transaction, you send that reversal. If you receive a response from the server, but the response comes late and you have already timed out, you also send a ‘late-response’ reversal to the server.

> In the same way, when you post a transaction that already took place (i.e. adjusting a previously approved transaction for a different amount, something that happens all the time at restaurants that support tips, or gas pump transactions where you approve for $100 but then complete for $20), and you don’t get a response, you send a retransmission of the same transaction, as many times as necessary in order to deal with transient network problems.

> In order to support reversals and retransmissions, you need a transaction unique identifier. Different networks use different transaction identifiers, either the Terminal ID + a Serial Trace Audit Number, or a Retrieval Reference Number, or in ISO-8583 v2003 the new Transaction life cycle identification data, the client generates a unique identifier so that when you send a follow-up transaction, you can send the server a reference to the original one.

__RRN (Retrieval Reference Number)__

* Client code could generate a UUID

__Support reversals (DELETE)__

* Could be as simple as adding a new verb, DELETE, where the only parameter can be — along with authentication data — the RRN

__Support for retransmissions (PUT)__

* If your initial transaction was a POST, I propose that you also accept a PUT, with the same parameters. On the server side the difference between the POST and the PUT operations would be just an extra step to make sure that you didn’t process the original POST, and return the original auth code if that was the case.
