[![Build Status](https://travis-ci.org/tomsanbear/recorder.svg?branch=master)](https://travis-ci.org/tomsanbear/recorder) [![codecov](https://codecov.io/gh/tomsanbear/recorder/branch/master/graph/badge.svg)](https://codecov.io/gh/tomsanbear/recorder) [![GitHub version](https://badge.fury.io/gh/tomsanbear%2Frecorder.svg)](https://badge.fury.io/gh/tomsanbear%2Frecorder) ![GitHub issues](https://img.shields.io/github/issues/tomsanbear/recorder.svg) ![GitHub pull requests](https://img.shields.io/github/issues-pr/tomsanbear/recorder.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/tomsanbear/recorder)](https://goreportcard.com/report/github.com/tomsanbear/recorder)


# CoreDNS recorder

This CoreDNS plugin implements a dns.ResponseWriter that records the message that was sent through it's chain. This is useful only if you need to access readonly data for what was sent back to the user. 

There is no configuration for this plugin at this time, super simple and easy implementation to add for your plugins.

For reference use, please see my other pluging [dns logging](https://github.com/tomsanbear/dnslogging)
