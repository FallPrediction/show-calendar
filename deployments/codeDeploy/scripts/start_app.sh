#! /bin/bash
systemctl enable souflair.service
systemctl daemon-reload
systemctl restart souflair.service
