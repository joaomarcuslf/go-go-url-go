#!/bin/bash

if [ ! -f .env ]
then
	cp sample.env .env
fi
