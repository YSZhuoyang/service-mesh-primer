#!/bin/bash

protoc -I ./googleapis -I. --include_imports --include_source_info --descriptor_set_out=greet.pb *.proto
