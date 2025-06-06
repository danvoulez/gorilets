#!/bin/bash
# Migration script for Minicontratos 2030
psql $DATABASE_URL -f ../../migrations/202501010001_create_users.sql
# Add more migration files as needed
