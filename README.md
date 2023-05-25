# VATPRC QueueMaster
## What
VATPRC QueueMaster is aiming to provide VATPRC airport controllers with a reliable source to manage their aircraft status and thereafter to **queue** their aircrafts, especially in an extremely busy online event. This repository only stores code for server-side and browser-side. For Euroscope plug-in, see [Ericple/VATPRC-UniSequence](https://github.com/Ericple/VATPRC-UniSequence).

## Why
> What is my sequence number?

In some large events, airport controlling positions usually become overwhelmed because lots of aircraft will have to be asked to stand by for pushback or delivery due to flow control for a long time. Usually, the controller will maintain an aircraft queue using pens and papers, or using the scratch pad on their controller client. However, using pens and paper is somewhat tedious, and data on the scratch pad can be lost in some situations. What makes the situation even worse is online pilots usually tend to be impatient and will start to ask for their sequence, although it is totally understandable - nobody likes to wait. However, radio frequency resources will be consistently occupied for questions asking for sequence and estimated time, causing further congestion. But this may be avoided by providing a system synchronising the controller software and allowing pilots to see their sequences. And that is exactly why this project is raised. 

## How
(A detailed explanation on how the system works is TBA.)

## API Reference
See [GitHub Wiki](https://github.com/websterzh/vatprc-queue/wiki/API-Reference).