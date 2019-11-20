# Event exporter

This controller watches cluster events, filters them according to the parameters passed to it and writes them to STDOUT.

In this way, the events can be picked up by log aggregators such as Fluentd.
