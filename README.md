# KMS En/Decrypt

This is sample encrypt or decrypt text by using AWS KMS key. You have to create custom key on the AWS first.

### Usage

To decrypt message

```terminal
❯ ./kms_deencryption decrypt --help
Usage of decrypt:
  -profile string
    	aws credential profile. (optional) (default "default")
  -text string
    	text to decrypt (Required)
```

To encrypt message

```terminal
❯ ./kms_deencryption encrypt --help
Usage of decrypt:
  -keyname string
        Key name that use for encrypt. (Required)
  -profile string
    	aws credential profile. (optional) (default "default")
  -text string
    	text to encrypt. (Required)
```

### Example

Encryption

```terminal
❯ ./kms_deencrypt encrypt --keyname LoggingKey --profile dev-plugin --text HelloWorld
=====================
NP+BAwEBB3BheWxvYWQB/4IAAQMBA0tleQEKAAEFTm9uY2UB/4QAAQdNZXNzYWdlAQoAAAAZ/4MBAQEJWzI0XXVpbnQ4Af+EAAEGATAAAP/r/4IB/6gBAgMAeM012noRUNv9Slg1MqAUOvaR5yeVwEhfKkKso3849i9VAUiZftZyav6iWa/QX9F2To4AAABuMGwGCSqGSIb3DQEHBqBfMF0CAQAwWAYJKoZIhvcNAQcBMB4GCWCGSAFlAwQBLjARBAy26MKW+7iARVKWHgkCARCAK0VZzFqfe3yuH/EoS2uXvHDokjgde37Yw9QDrlT+VN+UO3yepVFsOp/jeO4BGFYLBChLc3coKf/KZBs3/61nVv+ILxpS/7//kf/w/44BGvtrAcLOriu+zH+hENWjegwgF2eJuZ05n2ISAA==
```

Decryption

```terminal
❯ ./kms_deencrypt decrypt --profile dev-plugin --text NP+BAwEBB3BheWxvYWQB/4IAAQMBA0tleQEKAAEFTm9uY2UB/4QAAQdNZXNzYWdlAQoAAAAZ/4MBAQEJWzI0XXVpbnQ4Af+EAAEGATAAAP/r/4IB/6gBAgMAeM012noRUNv9Slg1MqAUOvaR5yeVwEhfKkKso3849i9VAUiZftZyav6iWa/QX9F2To4AAABuMGwGCSqGSIb3DQEHBqBfMF0CAQAwWAYJKoZIhvcNAQcBMB4GCWCGSAFlAwQBLjARBAy26MKW+7iARVKWHgkCARCAK0VZzFqfe3yuH/EoS2uXvHDokjgde37Yw9QDrlT+VN+UO3yepVFsOp/jeO4BGFYLBChLc3coKf/KZBs3/61nVv+ILxpS/7//kf/w/44BGvtrAcLOriu+zH+hENWjegwgF2eJuZ05n2ISAA==
=====================
HelloWorld
```

That's it, peace :v:
