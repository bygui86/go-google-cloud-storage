
# GCP/GCS permissions

ℹ️ permission assigned to `service account` at `bucket level`.

## Legend

| Symbol | Description                   |
|--------|-------------------------------|
| ✅      | permission granted            |
| ❌      | permission denied (error 403) |
| ❓      | not tested                    |
| ❗️     | TBD                           |

## Bucket level

| Role                         | Bucket | list | exist | info | create | delete |
|------------------------------|--------|------|-------|------|--------|--------|
| Storage Admin                | same   | ❌    | ✅     | ✅    | ❓      | ❓      |
| Storage Admin                | others | ❌    | ❌     | ❌    | ❓      | ❓      |
| Storage Object Admin         | same   |      |       |      | ❓      | ❓      |
| Storage Object Admin         | others |      |       |      | ❓      | ❓      |
| Storage Object Creator       | same   |      |       |      | ❓      | ❓      |
| Storage Object Creator       | others |      |       |      | ❓      | ❓      |
| Storage Object Viewer        | same   |      |       |      | ❓      | ❓      |
| Storage Object Viewer        | others |      |       |      | ❓      | ❓      |
| Storage Legacy Bucket Owner  | same   |      |       |      | ❓      | ❓      |
| Storage Legacy Bucket Owner  | others |      |       |      | ❓      | ❓      |
| Storage Legacy Bucket Reader | same   |      |       |      | ❓      | ❓      |
| Storage Legacy Bucket Reader | others |      |       |      | ❓      | ❓      |
| Storage Legacy Bucket Writer | same   |      |       |      | ❓      | ❓      |
| Storage Legacy Bucket Writer | others |      |       |      | ❓      | ❓      |

## Object level

| Role          | Bucket | list | exist | info | upload | download | delete |
|---------------|--------|------|-------|------|--------|----------|--------|
| Storage Admin | same   | ✅    | ✅     | ✅    | ❓      | ✅        | ❓      |
| Storage Admin | others | ❌    | ❓     | ❓    | ❓      | ❓        | ❓      |
| Storage Object Admin         | same     |       |      |        |          |        |
| Storage Object Admin         | others     |       |      |        |          |        |
| Storage Object Creator       | same     |       |      |        |          |        |
| Storage Object Creator       | others     |       |      |        |          |        |
| Storage Object Viewer        | same     |       |      |        |          |        |
| Storage Object Viewer        | others     |       |      |        |          |        |
| Storage Legacy Object Owner  | same     |       |      |        |          |        |
| Storage Legacy Object Owner  | others     |       |      |        |          |        |
| Storage Legacy Object Reader | same     |       |      |        |          |        |
| Storage Legacy Object Reader | others     |       |      |        |          |        |
