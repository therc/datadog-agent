# Unless explicitly stated otherwise all files in this repository are licensed
# under the Apache License Version 2.0.
# This product includes software developed at Datadog (https:#www.datadoghq.com/).
# Copyright 2018 Datadog, Inc.

windows_arch :x86_64
# Don't append a timestamp to the package version
append_timestamp false

if ENV["S3_OMNIBUS_CACHE_BUCKET"]
  use_s3_caching true
  s3_bucket ENV["S3_OMNIBUS_CACHE_BUCKET"]
  s3_access_key ENV["AWS_ACCESS_KEY_ID"]
  s3_secret_key ENV["AWS_SECRET_ACCESS_KEY"]
end
