# alert-sink-go


### Basic Design Flow
POST /alerts
    |
    v
[ACCEPT] --decode--> events -> (chan Event) ---> [PROCESS]
                                            validate → route → batch --> (chan Batch)
                                                                     |
                                                                     v
                                                                 [DELIVER] -- Sender.Send