#!/bin/bash
slot=ai_platform_channel_group
psql -U postgres -c "SELECT pg_drop_replication_slot('"$slot"');"