#!/bin/bash
# Rollback script for Minicontratos 2030
kubectl rollout undo deployment/minicontratos-backend
