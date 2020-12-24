package gh

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// parse errors
var (
	ErrEventNotSpecifiedToParse      = errors.New("no Event specified to parse")
	ErrInvalidHTTPMethod             = errors.New("invalid HTTP Method")
	ErrMissingGitLabEventHeader      = errors.New("missing X-Gitlab-Event Header")
	ErrGitLabTokenVerificationFailed = errors.New("X-Gitlab-Token validation failed")
	ErrEventNotFound                 = errors.New("event not defined to be parsed")
	ErrParsingPayload                = errors.New("error parsing payload")
	ErrParsingSystemPayload          = errors.New("error parsing system payload")
	// ErrHMACVerificationFailed    = errors.New("HMAC verification failed")
)

// GitLab hook types
const (
	PushEvents               Event = "Push Hook"
	TagPushEvents                Event = "Tag Push Hook"
	Comments            Event = "Note Hook"
	// ConfidentialComments Event = "Confidential Comments Hook"
	IssuesEvents             Event = "Issue Hook"
	ConfidentialIssuesEvents Event = "Confidential Issue Hook"
	MergeRequestEvents       Event = "Merge Request Hook"
	JobEvents                Event = "Job Hook"
	PipelineEvents           Event = "Pipeline Hook"
	// WikiPageEvents           Event = "Wiki Page Hook"
	// DeploymentEvents Event = "Deployment Hook"
    // FeatureFlagevents Event = "Feature Flag events"
	ReleaseseEvents Event = "Release Hook"

	SystemHookEvents         Event = "System Hook"

	systemGroupCreate         string = "group_create"
	systemGroupDestroy         string = "group_destroy"
	systemGroupRename         string = "group_rename"
	systemKeyCreate         string = "key_create"
	systemKeyDestroy         string = "key_destroy"
	systemProjectCreate         string = "project_create"
	systemProjectDestroy         string = "project_destroy"
	systemProjectRename         string = "project_rename"
	systemProjectTransfer         string = "project_transfer"
	systemProjectUpdate         string = "project_update"
	systemUserAddToGroup         string = "user_add_to_group"
	systemUserAddToTeam         string = "user_add_to_team"
	systemUserCreate         string = "user_create"
	systemUserDestroy         string = "user_destroy"
	systemUserFailedLogin         string = "user_failed_login"
	systemUserRemoveFromGroup         string = "user_remove_from_group"
	systemUserRemoveFromTeam         string = "user_remove_from_team"
	systemUserRename         string = "user_rename"
	systemUserUpdateForGroup         string = "user_update_for_group"
	systemUserUpdateForTeam         string = "user_update_for_team"
	systemPush string = "push"
	systemTagPush string = "tag_push"
	systemMergeRequest string = "merge_request"
	systemRepositoryUpdate string = "repository_update"
)

// Option is a configuration option for the webhook
type Option func(*Webhook) error

// Options is a namespace var for configuration options
var Options = WebhookOptions{}

// WebhookOptions is a namespace for configuration option methods
type WebhookOptions struct{}

// Secret registers the GitLab secret
func (WebhookOptions) Secret(secret string) Option {
	return func(hook *Webhook) error {
		hook.secret = secret
		return nil
	}
}

// Webhook instance contains all methods needed to process events
type Webhook struct {
	secret string
}

// Event defines a GitLab hook event type by the X-Gitlab-Event Header
type Event string

// New creates and returns a WebHook instance denoted by the Provider type
func New(options ...Option) (*Webhook, error) {
	hook := new(Webhook)
	for _, opt := range options {
		if err := opt(hook); err != nil {
			return nil, errors.New("Error applying Option")
		}
	}
	return hook, nil
}

// Parse verifies and parses the events specified and returns the payload object or an error
func (hook Webhook) Parse(r *http.Request, events ...Event) (interface{}, error) {
	defer func() {
		_, _ = io.Copy(ioutil.Discard, r.Body)
		_ = r.Body.Close()
	}()

	if len(events) == 0 {
		return nil, ErrEventNotSpecifiedToParse
	}
	if r.Method != http.MethodPost {
		return nil, ErrInvalidHTTPMethod
	}

	// If we have a Secret set, we should check the MAC
	if len(hook.secret) > 0 {
		signature := r.Header.Get("X-Gitlab-Token")
		if signature != hook.secret {
			return nil, ErrGitLabTokenVerificationFailed
		}
	}

	event := r.Header.Get("X-Gitlab-Event")
	if len(event) == 0 {
		return nil, ErrMissingGitLabEventHeader
	}

	gitLabEvent := Event(event)

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return nil, ErrParsingPayload
	}

	return eventParsing(gitLabEvent, events, payload)
}

func eventParsing(gitLabEvent Event, events []Event, payload []byte) (interface{}, error) {

	var found bool
	for _, evt := range events {
		if evt == gitLabEvent {
			found = true
			break
		}
	}
	// event not defined to be parsed
	if !found {
		return nil, ErrEventNotFound
	}

	switch gitLabEvent {
	case PushEvents:  // 2002
		var pl PushEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case TagPushEvents: //2020
		var pl TagPushEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case IssuesEvents,ConfidentialIssuesEvents: // 2020
		var pl IssueEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case Comments:
		var pl CommentEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case MergeRequestEvents:
		var pl MergeRequestEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case PipelineEvents:
		var pl PipelineEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case ReleaseseEvents:
		var pl ReleaseseEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err
	case JobEvents:
		var pl JobEventPayload
		err := json.Unmarshal([]byte(payload), &pl)
		return pl, err

	case SystemHookEvents:
		var pl SystemHookPayload
		err := json.Unmarshal([]byte(payload), &pl)
		if err != nil {
			return nil, err
		}
		switch pl.EventName {
		case systemProjectCreate: // Project created
			var sl SystemProjectCreatePayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemProjectDestroy: // Project destroyed
			var sl SystemProjectDestroyPayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemProjectRename: // Project renamed Note that project_rename is not triggered if the namespace changes. Please refer to group_rename and user_rename for that case.
			var sl SystemProjectRenamePayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemProjectTransfer:
			var sl SystemProjectTransferPayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemProjectUpdate:
			var sl SystemProjectUpdatePayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemUserAddToTeam: // New Team Member
			var sl SystemUserAddToTeamPayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemUserRemoveFromTeam: // Team Member Removed
			var sl SystemUserRemoveFromTeamPayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemUserUpdateForTeam: // Team Member Updated:
			var sl SystemUserUpdateForTeamPayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemUserCreate:
			var sl SystemUserCreatePayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemUserDestroy:
			var sl SystemUserDestroyPayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemUserFailedLogin: // User failed login If the user is blocked via LDAP, state will be ldap_blocked.
			var sl SystemUserFailedLoginPayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemUserRename:
			var sl SystemUserRenamePayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemKeyCreate:
			var sl SystemKeyCreatePayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemKeyDestroy:
			var sl SystemKeyDestroyPayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemGroupCreate, systemGroupDestroy: // Group created owner_name and owner_email are always null
			var sl SystemGroupCreatePayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemGroupRename:
			var sl SystemGroupRenamePayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemUserAddToGroup, systemUserRemoveFromGroup, systemUserUpdateForGroup: // New Group Member
			var sl SystemUserGroupPayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemPush: // Push events Triggered when you push to the repository, except when pushing tags. It generates one event per modified branch.
			var sl SystemPushPayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemTagPush: // Triggered when you create (or delete) tags to the repository. It generates one event per modified tag.
			var sl SystemTagPushPayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemMergeRequest: // Triggered when a new merge request is created, an existing merge request was updated/merged/closed or a commit is added in the source branch.
			var sl SystemMergeRequestPayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		case systemRepositoryUpdate:
			var sl SystemRepositoryUpdatePayload
			err := json.Unmarshal([]byte(payload), &sl)
			return sl, err
		default:
			return nil, fmt.Errorf("unknown system hook event %s", gitLabEvent)
		}
	default:
		return nil, fmt.Errorf("unknown event %s", gitLabEvent)
	}
}