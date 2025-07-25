# gtasks CLI

A command-line interface (CLI) for managing your Google Tasks.

## Table of Contents

- [1. Authentication](#1-authentication)
- [2. Command Structure](#2-command-structure)
- [3. Offline Mode](#3-offline-mode)
- [4. Terminology](#4-terminology)
- [5. Command Reference](#5-command-reference)
  - [Account Management](#account-management)
  - [TaskList Management](#tasklist-management)
  - [Task Management](#task-management)
- [6. Implementation Details](#6-implementation-details)
- [7. Project Documentation](#7-project-documentation)
- [8. Running Tests](#8-running-tests)

---

## 1. Authentication

- **Google Sign-In:** The CLI authenticates with Google using OAuth 2.0.
- **Credential Caching:** Caches credentials locally for automatic use until they expire.
- **Token Refresh:** Automatically refreshes expired tokens.
- **Multi-Account Support:** Manage multiple Google accounts seamlessly.

## 2. Command Structure

The CLI follows a `gtasks <resource> <action> [flags]` pattern.

- **`resource`**: The type of object to operate on (e.g., `accounts`, `tasklists`, `tasks`).
- **`action`**: The operation to perform (e.g., `list`, `create`, `get`, `update`, `delete`).

## 3. Offline Mode

`gtasks` supports a full offline mode. By using the global `--offline` flag, you can manage your tasks and task lists without an internet connection. All changes are saved to a local file (`~/.config/gtasks/offline.json`).

**Note:** Synchronization must be handled manually. This tool does not currently provide a `sync` command.

---

## 4. Terminology

- **Account:** Refers to the Google Account you authenticate with via the SSO sign-in flow. The CLI can cache multiple accounts, but only one is active at a time.
- **TaskList:** A container for your tasks. A user can have multiple task lists to organize different areas of their life (e.g., "Work," "Groceries," "Personal Projects"). Each task list has a unique ID.
- **Task:** A single to-do item that exists within a specific TaskList. It has properties like a title, notes, due date, and a completion status. Each task has a unique ID.

---

## 5. Command Reference

### Account Management

Manage your authenticated Google accounts.

#### `gtasks login`
Initiates the Google SSO flow to authenticate a new user. The new account becomes the active one.
- **Usage:** `gtasks login`

#### `gtasks logout`
Removes the cached credentials for the currently active user.
- **Usage:** `gtasks logout`

#### `gtasks accounts list`
Lists all authenticated Google accounts.
- **Usage:** `gtasks accounts list`

#### `gtasks accounts switch`
Switches the active user to another authenticated account.
- **Usage:** `gtasks accounts switch <email>`
- **Arguments:**
  - `<email>` (required): The email address of the account to make active.

---

### TaskList Management

Manage your task lists.

#### `gtasks tasklists list`
Lists all task lists.
- **Usage:** `gtasks tasklists list`

#### `gtasks tasklists get`
Retrieves the details of a specific task list.
- **Usage:** `gtasks tasklists get <tasklist_id>`
- **Arguments:**
  - `<tasklist_id>` (required): The ID of the task list to retrieve.

#### `gtasks tasklists create`
Creates a new task list.
- **Usage:** `gtasks tasklists create --title <list_title>`
- **Flags:**
  - `--title` (string, required): The title for the new task list.

#### `gtasks tasklists update`
Updates the title of an existing task list.
- **Usage:** `gtasks tasklists update <tasklist_id> --title <new_title>`
- **Arguments:**
  - `<tasklist_id>` (required): The ID of the task list to update.
- **Flags:**
  - `--title` (string, required): The new title for the task list.

#### `gtasks tasklists delete`
Permanently deletes a task list and all of its tasks.
- **Usage:** `gtasks tasklists delete <tasklist_id>`
- **Arguments:**
  - `<tasklist_id>` (required): The ID of the task list to delete.

---

### Task Management

Manage your tasks within a task list.

#### `gtasks tasks list`
Lists tasks within a specific task list.
- **Usage:** `gtasks tasks list [--tasklist <tasklist_id>]`
- **Flags:**
  - `--tasklist` (string, optional): The ID of the task list. Defaults to `@default`.
  - `--show-completed` (boolean, optional): Include completed tasks.
  - `--show-hidden` (boolean, optional): Include hidden tasks.

#### `gtasks tasks get`
Retrieves the details of a specific task.
- **Usage:** `gtasks tasks get <task_id> [--tasklist <tasklist_id>]`
- **Arguments:**
  - `<task_id>` (required): The ID of the task.
- **Flags:**
  - `--tasklist` (string, optional): The ID of the task list containing the task. Defaults to `@default`.

#### `gtasks tasks create`
Creates a new task in a task list.
- **Usage:** `gtasks tasks create --title <task_title> [--tasklist <tasklist_id>]`
- **Flags:**
  - `--tasklist` (string, optional): The ID of the task list. Defaults to `@default`.
  - `--title` (string, required): The title of the task.
  - `--notes` (string, optional): Notes or description for the task.
  - `--due` (string, optional): Due date in RFC3339 format (e.g., "2025-12-31T22:00:00.000Z").

#### `gtasks tasks update`
Updates an existing task.
- **Usage:** `gtasks tasks update <task_id> [--tasklist <tasklist_id>] [flags]`
- **Arguments:**
  - `<task_id>` (required): The ID of the task to update.
- **Flags:**
  - `--tasklist` (string, optional): The ID of the task list. Defaults to `@default`.
  - `--title` (string, optional): The new title for the task.
  - `--notes` (string, optional): The new notes for the task.
  - `--due` (string, optional): The new due date in RFC3339 format.

#### `gtasks tasks complete`
Marks a task as complete.
- **Usage:** `gtasks tasks complete <task_id> [--tasklist <tasklist_id>]`
- **Arguments:**
  - `<task_id>` (required): The ID of the task.
- **Flags:**
  - `--tasklist` (string, optional): The ID of the task list. Defaults to `@default`.

#### `gtasks tasks uncomplete`
Marks a task as not complete.
- **Usage:** `gtasks tasks uncomplete <task_id> [--tasklist <tasklist_id>]`
- **Arguments:**
  - `<task_id>` (required): The ID of the task.
- **Flags:**
  - `--tasklist` (string, optional): The ID of the task list. Defaults to `@default`.

#### `gtasks tasks delete`
Permanently deletes a task.
- **Usage:** `gtasks tasks delete <task_id> [--tasklist <tasklist_id>]`
- **Arguments:**
  - `<task_id>` (required): The ID of the task.
- **Flags:**
  - `--tasklist` (string, optional): The ID of the task list. Defaults to `@default`.

## 6. Implementation Details

- **Language:** Go
- **Libraries:**
  - Cobra (`github.com/spf13/cobra`) for CLI structure.
  - Google API Client for Go (`google.golang.org/api/tasks/v1`).
  - Go OAuth2 Library (`golang.org/x/oauth2`).

## 7. Project Documentation

For more detailed information on the design and implementation, see the following documents:

- [Software Design (`DESIGN.md`)](./DESIGN.md)
- [Implementation Plan (`IMPLEMENTATION_PLAN.md`)](./IMPLEMENTATION_PLAN.md)
- [Worklog (`WORKLOG.md`)](./WORKLOG.md)
- [OAuth Web Flow Plan (`OAUTH_WEB_FLOW_PLAN.md`)](./OAUTH_WEB_FLOW_PLAN.md)

## 8. Running Tests

To run the full suite of tests, navigate to the `go/` directory and use the following command:

```bash
go test ./...
```

This command executes all unit and integration tests against a high-fidelity, in-memory mock of the Google Tasks API, ensuring that no real network calls are made and no authentication is required. It also runs the basic E2E tests.
