package gh

import (
	"time"
)

type PushEventPayload struct {
	ObjectKind  string `json:"object_kind"`
	EventName   string `json:"event_name"`
	Before      string `json:"before"`
	After       string `json:"after"`
	Ref         string `json:"ref"`
	CheckoutSha string `json:"checkout_sha"`
	Message     string `json:"message"`
	UserID      int    `json:"user_id"`
	UserName    string `json:"user_name"`
	UserEmail   string `json:"user_email"`
	UserAvatar  string `json:"user_avatar"`
	ProjectID   int    `json:"project_id"`
	Project     struct {
		ID                int    `json:"id"`
		Name              string `json:"name"`
		Description       string `json:"description"`
		WebURL            string `json:"web_url"`
		AvatarURL         string `json:"avatar_url"`
		GitSSHURL         string `json:"git_ssh_url"`
		GitHTTPURL        string `json:"git_http_url"`
		Namespace         string `json:"namespace"`
		VisibilityLevel   int    `json:"visibility_level"`
		PathWithNamespace string `json:"path_with_namespace"`
		DefaultBranch     string `json:"default_branch"`
	} `json:"project"`
	Commits []struct {
		ID        string    `json:"id"`
		Message   string    `json:"message"`
		Title     string    `json:"title"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
	} `json:"commits"`
	TotalCommitsCount int `json:"total_commits_count"`
	PushOptions       struct {
		Ci struct {
			Skip bool `json:"skip"`
		} `json:"ci"`
	} `json:"push_options"`
}

type TagPushEventPayload struct {
	ObjectKind  string `json:"object_kind"`
	EventName   string `json:"event_name"`
	Before      string `json:"before"`
	After       string `json:"after"`
	Ref         string `json:"ref"`
	CheckoutSha string `json:"checkout_sha"`
	Message     string `json:"message"`
	UserID      int    `json:"user_id"`
	UserName    string `json:"user_name"`
	UserEmail   string `json:"user_email"`
	UserAvatar  string `json:"user_avatar"`
	ProjectID   int    `json:"project_id"`
	Project     struct {
		ID                int    `json:"id"`
		Name              string `json:"name"`
		Description       string `json:"description"`
		WebURL            string `json:"web_url"`
		AvatarURL         string `json:"avatar_url"`
		GitSSHURL         string `json:"git_ssh_url"`
		GitHTTPURL        string `json:"git_http_url"`
		Namespace         string `json:"namespace"`
		VisibilityLevel   int    `json:"visibility_level"`
		PathWithNamespace string `json:"path_with_namespace"`
		DefaultBranch     string `json:"default_branch"`
	} `json:"project"`
	Commits []struct {
		ID        string    `json:"id"`
		Message   string    `json:"message"`
		Title     string    `json:"title"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
	} `json:"commits"`
	TotalCommitsCount int `json:"total_commits_count"`
	PushOptions       struct {
		Ci struct {
			Skip bool `json:"skip"`
		} `json:"ci"`
	} `json:"push_options"`
}

type IssueEventPayload struct {
	ObjectKind string `json:"object_kind"`
	EventType  string `json:"event_type"`
	User       struct {
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	Project struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebURL            string      `json:"web_url"`
		AvatarURL         interface{} `json:"avatar_url"`
		GitSSHURL         string      `json:"git_ssh_url"`
		GitHTTPURL        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		CiConfigPath      interface{} `json:"ci_config_path"`
		Homepage          string      `json:"homepage"`
		URL               string      `json:"url"`
		SSHURL            string      `json:"ssh_url"`
		HTTPURL           string      `json:"http_url"`
	} `json:"project"`
	ObjectAttributes struct {
		AuthorID            int           `json:"author_id"`
		ClosedAt            interface{}   `json:"closed_at"`
		Confidential        bool          `json:"confidential"`
		CreatedAt           string        `json:"created_at"`
		Description         string        `json:"description"`
		DiscussionLocked    interface{}   `json:"discussion_locked"`
		DueDate             interface{}   `json:"due_date"`
		ID                  int           `json:"id"`
		Iid                 int           `json:"iid"`
		LastEditedAt        interface{}   `json:"last_edited_at"`
		LastEditedByID      interface{}   `json:"last_edited_by_id"`
		MilestoneID         interface{}   `json:"milestone_id"`
		MovedToID           interface{}   `json:"moved_to_id"`
		DuplicatedToID      interface{}   `json:"duplicated_to_id"`
		ProjectID           int           `json:"project_id"`
		RelativePosition    int           `json:"relative_position"`
		StateID             int           `json:"state_id"`
		TimeEstimate        int           `json:"time_estimate"`
		Title               string        `json:"title"`
		UpdatedAt           string        `json:"updated_at"`
		UpdatedByID         interface{}   `json:"updated_by_id"`
		Weight              interface{}   `json:"weight"`
		URL                 string        `json:"url"`
		TotalTimeSpent      int           `json:"total_time_spent"`
		HumanTotalTimeSpent interface{}   `json:"human_total_time_spent"`
		HumanTimeEstimate   interface{}   `json:"human_time_estimate"`
		AssigneeIds         []interface{} `json:"assignee_ids"`
		AssigneeID          interface{}   `json:"assignee_id"`
		Labels              []interface{} `json:"labels"`
		State               string        `json:"state"`
	} `json:"object_attributes"`
	Labels  []interface{} `json:"labels"`
	Changes struct {
	} `json:"changes"`
	Repository struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
	} `json:"repository"`
}

type CommentEventPayload struct {
	ObjectKind string `json:"object_kind"`
	EventType  string `json:"event_type"`
	User       struct {
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	ProjectID int `json:"project_id"`
	Project   struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebURL            string      `json:"web_url"`
		AvatarURL         interface{} `json:"avatar_url"`
		GitSSHURL         string      `json:"git_ssh_url"`
		GitHTTPURL        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		CiConfigPath      interface{} `json:"ci_config_path"`
		Homepage          string      `json:"homepage"`
		URL               string      `json:"url"`
		SSHURL            string      `json:"ssh_url"`
		HTTPURL           string      `json:"http_url"`
	} `json:"project"`
	ObjectAttributes struct {
		Attachment       interface{} `json:"attachment"`
		AuthorID         int         `json:"author_id"`
		ChangePosition   interface{} `json:"change_position"`
		CommitID         interface{} `json:"commit_id"`
		CreatedAt        string      `json:"created_at"`
		DiscussionID     string      `json:"discussion_id"`
		ID               int         `json:"id"`
		LineCode         interface{} `json:"line_code"`
		Note             string      `json:"note"`
		NoteableID       int         `json:"noteable_id"`
		NoteableType     string      `json:"noteable_type"`
		OriginalPosition interface{} `json:"original_position"`
		Position         interface{} `json:"position"`
		ProjectID        int         `json:"project_id"`
		ResolvedAt       interface{} `json:"resolved_at"`
		ResolvedByID     interface{} `json:"resolved_by_id"`
		ResolvedByPush   interface{} `json:"resolved_by_push"`
		StDiff           interface{} `json:"st_diff"`
		System           bool        `json:"system"`
		Type             interface{} `json:"type"`
		UpdatedAt        string      `json:"updated_at"`
		UpdatedByID      interface{} `json:"updated_by_id"`
		Description      string      `json:"description"`
		URL              string      `json:"url"`
	} `json:"object_attributes"`
	Repository struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
	} `json:"repository"`
	MergeRequest struct {
		AssigneeID     interface{} `json:"assignee_id"`
		AuthorID       int         `json:"author_id"`
		CreatedAt      string      `json:"created_at"`
		Description    string      `json:"description"`
		HeadPipelineID interface{} `json:"head_pipeline_id"`
		ID             int         `json:"id"`
		Iid            int         `json:"iid"`
		LastEditedAt   interface{} `json:"last_edited_at"`
		LastEditedByID interface{} `json:"last_edited_by_id"`
		MergeCommitSha interface{} `json:"merge_commit_sha"`
		MergeError     interface{} `json:"merge_error"`
		MergeParams    struct {
			ForceRemoveSourceBranch string `json:"force_remove_source_branch"`
		} `json:"merge_params"`
		MergeStatus               string      `json:"merge_status"`
		MergeUserID               interface{} `json:"merge_user_id"`
		MergeWhenPipelineSucceeds bool        `json:"merge_when_pipeline_succeeds"`
		MilestoneID               interface{} `json:"milestone_id"`
		SourceBranch              string      `json:"source_branch"`
		SourceProjectID           int         `json:"source_project_id"`
		StateID                   int         `json:"state_id"`
		TargetBranch              string      `json:"target_branch"`
		TargetProjectID           int         `json:"target_project_id"`
		TimeEstimate              int         `json:"time_estimate"`
		Title                     string      `json:"title"`
		UpdatedAt                 string      `json:"updated_at"`
		UpdatedByID               interface{} `json:"updated_by_id"`
		URL                       string      `json:"url"`
		Source                    struct {
			ID                int         `json:"id"`
			Name              string      `json:"name"`
			Description       string      `json:"description"`
			WebURL            string      `json:"web_url"`
			AvatarURL         interface{} `json:"avatar_url"`
			GitSSHURL         string      `json:"git_ssh_url"`
			GitHTTPURL        string      `json:"git_http_url"`
			Namespace         string      `json:"namespace"`
			VisibilityLevel   int         `json:"visibility_level"`
			PathWithNamespace string      `json:"path_with_namespace"`
			DefaultBranch     string      `json:"default_branch"`
			CiConfigPath      interface{} `json:"ci_config_path"`
			Homepage          string      `json:"homepage"`
			URL               string      `json:"url"`
			SSHURL            string      `json:"ssh_url"`
			HTTPURL           string      `json:"http_url"`
		} `json:"source"`
		Target struct {
			ID                int         `json:"id"`
			Name              string      `json:"name"`
			Description       string      `json:"description"`
			WebURL            string      `json:"web_url"`
			AvatarURL         interface{} `json:"avatar_url"`
			GitSSHURL         string      `json:"git_ssh_url"`
			GitHTTPURL        string      `json:"git_http_url"`
			Namespace         string      `json:"namespace"`
			VisibilityLevel   int         `json:"visibility_level"`
			PathWithNamespace string      `json:"path_with_namespace"`
			DefaultBranch     string      `json:"default_branch"`
			CiConfigPath      interface{} `json:"ci_config_path"`
			Homepage          string      `json:"homepage"`
			URL               string      `json:"url"`
			SSHURL            string      `json:"ssh_url"`
			HTTPURL           string      `json:"http_url"`
		} `json:"target"`
		LastCommit struct {
			ID        string    `json:"id"`
			Message   string    `json:"message"`
			Title     string    `json:"title"`
			Timestamp time.Time `json:"timestamp"`
			URL       string    `json:"url"`
			Author    struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
		} `json:"last_commit"`
		WorkInProgress      bool          `json:"work_in_progress"`
		TotalTimeSpent      int           `json:"total_time_spent"`
		HumanTotalTimeSpent interface{}   `json:"human_total_time_spent"`
		HumanTimeEstimate   interface{}   `json:"human_time_estimate"`
		AssigneeIds         []interface{} `json:"assignee_ids"`
		State               string        `json:"state"`
	} `json:"merge_request"`
}

type MergeRequestEventPayload struct {
	ObjectKind string `json:"object_kind"`
	EventType  string `json:"event_type"`
	User       struct {
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	Project struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebURL            string      `json:"web_url"`
		AvatarURL         interface{} `json:"avatar_url"`
		GitSSHURL         string      `json:"git_ssh_url"`
		GitHTTPURL        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		CiConfigPath      interface{} `json:"ci_config_path"`
		Homepage          string      `json:"homepage"`
		URL               string      `json:"url"`
		SSHURL            string      `json:"ssh_url"`
		HTTPURL           string      `json:"http_url"`
	} `json:"project"`
	ObjectAttributes struct {
		AssigneeID     interface{} `json:"assignee_id"`
		AuthorID       int         `json:"author_id"`
		CreatedAt      string      `json:"created_at"`
		Description    string      `json:"description"`
		HeadPipelineID interface{} `json:"head_pipeline_id"`
		ID             int         `json:"id"`
		Iid            int         `json:"iid"`
		LastEditedAt   interface{} `json:"last_edited_at"`
		LastEditedByID interface{} `json:"last_edited_by_id"`
		MergeCommitSha interface{} `json:"merge_commit_sha"`
		MergeError     interface{} `json:"merge_error"`
		MergeParams    struct {
			ForceRemoveSourceBranch string `json:"force_remove_source_branch"`
		} `json:"merge_params"`
		MergeStatus               string      `json:"merge_status"`
		MergeUserID               interface{} `json:"merge_user_id"`
		MergeWhenPipelineSucceeds bool        `json:"merge_when_pipeline_succeeds"`
		MilestoneID               interface{} `json:"milestone_id"`
		SourceBranch              string      `json:"source_branch"`
		SourceProjectID           int         `json:"source_project_id"`
		StateID                   int         `json:"state_id"`
		TargetBranch              string      `json:"target_branch"`
		TargetProjectID           int         `json:"target_project_id"`
		TimeEstimate              int         `json:"time_estimate"`
		Title                     string      `json:"title"`
		UpdatedAt                 string      `json:"updated_at"`
		UpdatedByID               interface{} `json:"updated_by_id"`
		URL                       string      `json:"url"`
		Source                    struct {
			ID                int         `json:"id"`
			Name              string      `json:"name"`
			Description       string      `json:"description"`
			WebURL            string      `json:"web_url"`
			AvatarURL         interface{} `json:"avatar_url"`
			GitSSHURL         string      `json:"git_ssh_url"`
			GitHTTPURL        string      `json:"git_http_url"`
			Namespace         string      `json:"namespace"`
			VisibilityLevel   int         `json:"visibility_level"`
			PathWithNamespace string      `json:"path_with_namespace"`
			DefaultBranch     string      `json:"default_branch"`
			CiConfigPath      interface{} `json:"ci_config_path"`
			Homepage          string      `json:"homepage"`
			URL               string      `json:"url"`
			SSHURL            string      `json:"ssh_url"`
			HTTPURL           string      `json:"http_url"`
		} `json:"source"`
		Target struct {
			ID                int         `json:"id"`
			Name              string      `json:"name"`
			Description       string      `json:"description"`
			WebURL            string      `json:"web_url"`
			AvatarURL         interface{} `json:"avatar_url"`
			GitSSHURL         string      `json:"git_ssh_url"`
			GitHTTPURL        string      `json:"git_http_url"`
			Namespace         string      `json:"namespace"`
			VisibilityLevel   int         `json:"visibility_level"`
			PathWithNamespace string      `json:"path_with_namespace"`
			DefaultBranch     string      `json:"default_branch"`
			CiConfigPath      interface{} `json:"ci_config_path"`
			Homepage          string      `json:"homepage"`
			URL               string      `json:"url"`
			SSHURL            string      `json:"ssh_url"`
			HTTPURL           string      `json:"http_url"`
		} `json:"target"`
		LastCommit struct {
			ID        string    `json:"id"`
			Message   string    `json:"message"`
			Title     string    `json:"title"`
			Timestamp time.Time `json:"timestamp"`
			URL       string    `json:"url"`
			Author    struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
		} `json:"last_commit"`
		WorkInProgress      bool          `json:"work_in_progress"`
		TotalTimeSpent      int           `json:"total_time_spent"`
		HumanTotalTimeSpent interface{}   `json:"human_total_time_spent"`
		HumanTimeEstimate   interface{}   `json:"human_time_estimate"`
		AssigneeIds         []interface{} `json:"assignee_ids"`
		State               string        `json:"state"`
	} `json:"object_attributes"`
	Labels  []interface{} `json:"labels"`
	Changes struct {
	} `json:"changes"`
	Repository struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
	} `json:"repository"`
}

type PipelineEventPayload struct {
	ObjectKind       string `json:"object_kind"`
	ObjectAttributes struct {
		ID             int           `json:"id"`
		Ref            string        `json:"ref"`
		Tag            bool          `json:"tag"`
		Sha            string        `json:"sha"`
		BeforeSha      string        `json:"before_sha"`
		Source         string        `json:"source"`
		Status         string        `json:"status"`
		DetailedStatus string        `json:"detailed_status"`
		Stages         []string      `json:"stages"`
		CreatedAt      string        `json:"created_at"`
		FinishedAt     interface{}   `json:"finished_at"`
		Duration       interface{}   `json:"duration"`
		Variables      []interface{} `json:"variables"`
	} `json:"object_attributes"`
	MergeRequest interface{} `json:"merge_request"`
	User         struct {
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	Project struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebURL            string      `json:"web_url"`
		AvatarURL         interface{} `json:"avatar_url"`
		GitSSHURL         string      `json:"git_ssh_url"`
		GitHTTPURL        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		CiConfigPath      interface{} `json:"ci_config_path"`
	} `json:"project"`
	Commit struct {
		ID        string    `json:"id"`
		Message   string    `json:"message"`
		Title     string    `json:"title"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
	} `json:"commit"`
	Builds []struct {
		ID           int         `json:"id"`
		Stage        string      `json:"stage"`
		Name         string      `json:"name"`
		Status       string      `json:"status"`
		CreatedAt    string      `json:"created_at"`
		StartedAt    interface{} `json:"started_at"`
		FinishedAt   interface{} `json:"finished_at"`
		When         string      `json:"when"`
		Manual       bool        `json:"manual"`
		AllowFailure bool        `json:"allow_failure"`
		User         struct {
			Name      string `json:"name"`
			Username  string `json:"username"`
			AvatarURL string `json:"avatar_url"`
			Email     string `json:"email"`
		} `json:"user"`
		Runner        interface{} `json:"runner"`
		ArtifactsFile struct {
			Filename interface{} `json:"filename"`
			Size     interface{} `json:"size"`
		} `json:"artifacts_file"`
	} `json:"builds"`
}

type JobEventPayload struct {
	ObjectKind         string  `json:"object_kind"`
	Ref                string  `json:"ref"`
	Tag                bool    `json:"tag"`
	BeforeSha          string  `json:"before_sha"`
	Sha                string  `json:"sha"`
	BuildID            int     `json:"build_id"`
	BuildName          string  `json:"build_name"`
	BuildStage         string  `json:"build_stage"`
	BuildStatus        string  `json:"build_status"`
	BuildStartedAt     string  `json:"build_started_at"`
	BuildFinishedAt    string  `json:"build_finished_at"`
	BuildDuration      float64 `json:"build_duration"`
	BuildAllowFailure  bool    `json:"build_allow_failure"`
	BuildFailureReason string  `json:"build_failure_reason"`
	PipelineID         int     `json:"pipeline_id"`
	Runner             struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
		Active      bool   `json:"active"`
		IsShared    bool   `json:"is_shared"`
	} `json:"runner"`
	ProjectID   int    `json:"project_id"`
	ProjectName string `json:"project_name"`
	User        struct {
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	Commit struct {
		ID          int    `json:"id"`
		Sha         string `json:"sha"`
		Message     string `json:"message"`
		AuthorName  string `json:"author_name"`
		AuthorEmail string `json:"author_email"`
		AuthorURL   string `json:"author_url"`
		Status      string `json:"status"`
		Duration    int    `json:"duration"`
		StartedAt   string `json:"started_at"`
		FinishedAt  string `json:"finished_at"`
	} `json:"commit"`
	Repository struct {
		Name            string `json:"name"`
		URL             string `json:"url"`
		Description     string `json:"description"`
		Homepage        string `json:"homepage"`
		GitHTTPURL      string `json:"git_http_url"`
		GitSSHURL       string `json:"git_ssh_url"`
		VisibilityLevel int    `json:"visibility_level"`
	} `json:"repository"`
}

type ReleaseseEventPayload struct {
	ID          int    `json:"id"`
	CreatedAt   string `json:"created_at"`
	Description string `json:"description"`
	Name        string `json:"name"`
	ReleasedAt  string `json:"released_at"`
	Tag         string `json:"tag"`
	ObjectKind  string `json:"object_kind"`
	Project     struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebURL            string      `json:"web_url"`
		AvatarURL         interface{} `json:"avatar_url"`
		GitSSHURL         string      `json:"git_ssh_url"`
		GitHTTPURL        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		CiConfigPath      interface{} `json:"ci_config_path"`
		Homepage          string      `json:"homepage"`
		URL               string      `json:"url"`
		SSHURL            string      `json:"ssh_url"`
		HTTPURL           string      `json:"http_url"`
	} `json:"project"`
	URL    string `json:"url"`
	Action string `json:"action"`
	Assets struct {
		Count   int           `json:"count"`
		Links   []interface{} `json:"links"`
		Sources []struct {
			Format string `json:"format"`
			URL    string `json:"url"`
		} `json:"sources"`
	} `json:"assets"`
	Commit struct {
		ID        string    `json:"id"`
		Message   string    `json:"message"`
		Title     string    `json:"title"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
	} `json:"commit"`
}

type SystemHookPayload struct {
	EventName  string `json:"event_name"`
}

type SystemProjectCreatePayload struct {
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	EventName         string    `json:"event_name"`
	Name              string    `json:"name"`
	OwnerEmail        string    `json:"owner_email"`
	OwnerName         string    `json:"owner_name"`
	Path              string    `json:"path"`
	PathWithNamespace string    `json:"path_with_namespace"`
	ProjectID         int       `json:"project_id"`
	ProjectVisibility string    `json:"project_visibility"`
}

type SystemProjectDestroyPayload struct {
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	EventName         string    `json:"event_name"`
	Name              string    `json:"name"`
	OwnerEmail        string    `json:"owner_email"`
	OwnerName         string    `json:"owner_name"`
	Path              string    `json:"path"`
	PathWithNamespace string    `json:"path_with_namespace"`
	ProjectID         int       `json:"project_id"`
	ProjectVisibility string    `json:"project_visibility"`
}

type SystemProjectRenamePayload struct {
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	EventName            string    `json:"event_name"`
	Name                 string    `json:"name"`
	Path                 string    `json:"path"`
	PathWithNamespace    string    `json:"path_with_namespace"`
	ProjectID            int       `json:"project_id"`
	OwnerName            string    `json:"owner_name"`
	OwnerEmail           string    `json:"owner_email"`
	ProjectVisibility    string    `json:"project_visibility"`
	OldPathWithNamespace string    `json:"old_path_with_namespace"`
}

type SystemProjectTransferPayload struct {
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	EventName            string    `json:"event_name"`
	Name                 string    `json:"name"`
	Path                 string    `json:"path"`
	PathWithNamespace    string    `json:"path_with_namespace"`
	ProjectID            int       `json:"project_id"`
	OwnerName            string    `json:"owner_name"`
	OwnerEmail           string    `json:"owner_email"`
	ProjectVisibility    string    `json:"project_visibility"`
	OldPathWithNamespace string    `json:"old_path_with_namespace"`
}

type SystemProjectUpdatePayload struct {
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	EventName         string    `json:"event_name"`
	Name              string    `json:"name"`
	OwnerEmail        string    `json:"owner_email"`
	OwnerName         string    `json:"owner_name"`
	Path              string    `json:"path"`
	PathWithNamespace string    `json:"path_with_namespace"`
	ProjectID         int       `json:"project_id"`
	ProjectVisibility string    `json:"project_visibility"`
}

type SystemUserAddToTeamPayload struct {
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	EventName                string    `json:"event_name"`
	AccessLevel              string    `json:"access_level"`
	ProjectID                int       `json:"project_id"`
	ProjectName              string    `json:"project_name"`
	ProjectPath              string    `json:"project_path"`
	ProjectPathWithNamespace string    `json:"project_path_with_namespace"`
	UserEmail                string    `json:"user_email"`
	UserName                 string    `json:"user_name"`
	UserUsername             string    `json:"user_username"`
	UserID                   int       `json:"user_id"`
	ProjectVisibility        string    `json:"project_visibility"`
}

type SystemUserRemoveFromTeamPayload struct {
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	EventName                string    `json:"event_name"`
	AccessLevel              string    `json:"access_level"`
	ProjectID                int       `json:"project_id"`
	ProjectName              string    `json:"project_name"`
	ProjectPath              string    `json:"project_path"`
	ProjectPathWithNamespace string    `json:"project_path_with_namespace"`
	UserEmail                string    `json:"user_email"`
	UserName                 string    `json:"user_name"`
	UserUsername             string    `json:"user_username"`
	UserID                   int       `json:"user_id"`
	ProjectVisibility        string    `json:"project_visibility"`
}

type SystemUserUpdateForTeamPayload struct {
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	EventName                string    `json:"event_name"`
	AccessLevel              string    `json:"access_level"`
	ProjectID                int       `json:"project_id"`
	ProjectName              string    `json:"project_name"`
	ProjectPath              string    `json:"project_path"`
	ProjectPathWithNamespace string    `json:"project_path_with_namespace"`
	UserEmail                string    `json:"user_email"`
	UserName                 string    `json:"user_name"`
	UserUsername             string    `json:"user_username"`
	UserID                   int       `json:"user_id"`
	ProjectVisibility        string    `json:"project_visibility"`
}

type SystemUserCreatePayload struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	EventName string    `json:"event_name"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	UserID    int       `json:"user_id"`
}

type SystemUserDestroyPayload struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	EventName string    `json:"event_name"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	UserID    int       `json:"user_id"`
}

type SystemUserFailedLoginPayload struct {
	EventName string    `json:"event_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	State     string    `json:"state"`
}

type SystemUserRenamePayload struct {
	EventName   string    `json:"event_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	UserID      int       `json:"user_id"`
	Username    string    `json:"username"`
	OldUsername string    `json:"old_username"`
}

type SystemKeyCreatePayload struct {
	EventName string    `json:"event_name"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Key       string    `json:"key"`
	ID        int       `json:"id"`
}

type SystemKeyDestroyPayload struct {
	EventName string    `json:"event_name"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Key       string    `json:"key"`
	ID        int       `json:"id"`
}

type SystemGroupCreatePayload struct {
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	EventName  string      `json:"event_name"`
	Name       string      `json:"name"`
	OwnerEmail interface{} `json:"owner_email"`
	OwnerName  interface{} `json:"owner_name"`
	Path       string      `json:"path"`
	GroupID    int         `json:"group_id"`
}

type SystemGroupRenamePayload struct {
	EventName   string      `json:"event_name"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	FullPath    string      `json:"full_path"`
	GroupID     int         `json:"group_id"`
	OwnerName   interface{} `json:"owner_name"`
	OwnerEmail  interface{} `json:"owner_email"`
	OldPath     string      `json:"old_path"`
	OldFullPath string      `json:"old_full_path"`
}

type SystemUserGroupPayload struct {
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	EventName    string    `json:"event_name"`
	GroupAccess  string    `json:"group_access"`
	GroupID      int       `json:"group_id"`
	GroupName    string    `json:"group_name"`
	GroupPath    string    `json:"group_path"`
	UserEmail    string    `json:"user_email"`
	UserName     string    `json:"user_name"`
	UserUsername string    `json:"user_username"`
	UserID       int       `json:"user_id"`
}

type SystemPushPayload struct {
	EventName   string `json:"event_name"`
	Before      string `json:"before"`
	After       string `json:"after"`
	Ref         string `json:"ref"`
	CheckoutSha string `json:"checkout_sha"`
	UserID      int    `json:"user_id"`
	UserName    string `json:"user_name"`
	UserEmail   string `json:"user_email"`
	UserAvatar  string `json:"user_avatar"`
	ProjectID   int    `json:"project_id"`
	Project     struct {
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebURL            string      `json:"web_url"`
		AvatarURL         interface{} `json:"avatar_url"`
		GitSSHURL         string      `json:"git_ssh_url"`
		GitHTTPURL        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		Homepage          string      `json:"homepage"`
		URL               string      `json:"url"`
		SSHURL            string      `json:"ssh_url"`
		HTTPURL           string      `json:"http_url"`
	} `json:"project"`
	Repository struct {
		Name            string `json:"name"`
		URL             string `json:"url"`
		Description     string `json:"description"`
		Homepage        string `json:"homepage"`
		GitHTTPURL      string `json:"git_http_url"`
		GitSSHURL       string `json:"git_ssh_url"`
		VisibilityLevel int    `json:"visibility_level"`
	} `json:"repository"`
	Commits []struct {
		ID        string    `json:"id"`
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
	} `json:"commits"`
	TotalCommitsCount int `json:"total_commits_count"`
}

type SystemTagPushPayload struct {
	EventName   string `json:"event_name"`
	Before      string `json:"before"`
	After       string `json:"after"`
	Ref         string `json:"ref"`
	CheckoutSha string `json:"checkout_sha"`
	UserID      int    `json:"user_id"`
	UserName    string `json:"user_name"`
	UserAvatar  string `json:"user_avatar"`
	ProjectID   int    `json:"project_id"`
	Project     struct {
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebURL            string      `json:"web_url"`
		AvatarURL         interface{} `json:"avatar_url"`
		GitSSHURL         string      `json:"git_ssh_url"`
		GitHTTPURL        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		Homepage          string      `json:"homepage"`
		URL               string      `json:"url"`
		SSHURL            string      `json:"ssh_url"`
		HTTPURL           string      `json:"http_url"`
	} `json:"project"`
	Repository struct {
		Name            string `json:"name"`
		URL             string `json:"url"`
		Description     string `json:"description"`
		Homepage        string `json:"homepage"`
		GitHTTPURL      string `json:"git_http_url"`
		GitSSHURL       string `json:"git_ssh_url"`
		VisibilityLevel int    `json:"visibility_level"`
	} `json:"repository"`
	Commits           []interface{} `json:"commits"`
	TotalCommitsCount int           `json:"total_commits_count"`
}

type SystemMergeRequestPayload struct {
	ObjectKind string `json:"object_kind"`
	User       struct {
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
	} `json:"user"`
	Project struct {
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebURL            string      `json:"web_url"`
		AvatarURL         interface{} `json:"avatar_url"`
		GitSSHURL         string      `json:"git_ssh_url"`
		GitHTTPURL        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		CiConfigPath      string      `json:"ci_config_path"`
		Homepage          string      `json:"homepage"`
		URL               string      `json:"url"`
		SSHURL            string      `json:"ssh_url"`
		HTTPURL           string      `json:"http_url"`
	} `json:"project"`
	ObjectAttributes struct {
		ID              int         `json:"id"`
		TargetBranch    string      `json:"target_branch"`
		SourceBranch    string      `json:"source_branch"`
		SourceProjectID int         `json:"source_project_id"`
		AuthorID        int         `json:"author_id"`
		AssigneeID      int         `json:"assignee_id"`
		Title           string      `json:"title"`
		CreatedAt       time.Time   `json:"created_at"`
		UpdatedAt       time.Time   `json:"updated_at"`
		MilestoneID     interface{} `json:"milestone_id"`
		State           string      `json:"state"`
		MergeStatus     string      `json:"merge_status"`
		TargetProjectID int         `json:"target_project_id"`
		Iid             int         `json:"iid"`
		Description     string      `json:"description"`
		UpdatedByID     int         `json:"updated_by_id"`
		MergeError      interface{} `json:"merge_error"`
		MergeParams     struct {
			ForceRemoveSourceBranch string `json:"force_remove_source_branch"`
		} `json:"merge_params"`
		MergeWhenPipelineSucceeds bool        `json:"merge_when_pipeline_succeeds"`
		MergeUserID               interface{} `json:"merge_user_id"`
		MergeCommitSha            interface{} `json:"merge_commit_sha"`
		DeletedAt                 interface{} `json:"deleted_at"`
		InProgressMergeCommitSha  interface{} `json:"in_progress_merge_commit_sha"`
		LockVersion               int         `json:"lock_version"`
		TimeEstimate              int         `json:"time_estimate"`
		LastEditedAt              time.Time   `json:"last_edited_at"`
		LastEditedByID            int         `json:"last_edited_by_id"`
		HeadPipelineID            int         `json:"head_pipeline_id"`
		RefFetched                bool        `json:"ref_fetched"`
		MergeJid                  interface{} `json:"merge_jid"`
		Source                    struct {
			Name              string      `json:"name"`
			Description       string      `json:"description"`
			WebURL            string      `json:"web_url"`
			AvatarURL         interface{} `json:"avatar_url"`
			GitSSHURL         string      `json:"git_ssh_url"`
			GitHTTPURL        string      `json:"git_http_url"`
			Namespace         string      `json:"namespace"`
			VisibilityLevel   int         `json:"visibility_level"`
			PathWithNamespace string      `json:"path_with_namespace"`
			DefaultBranch     string      `json:"default_branch"`
			CiConfigPath      string      `json:"ci_config_path"`
			Homepage          string      `json:"homepage"`
			URL               string      `json:"url"`
			SSHURL            string      `json:"ssh_url"`
			HTTPURL           string      `json:"http_url"`
		} `json:"source"`
		Target struct {
			Name              string      `json:"name"`
			Description       string      `json:"description"`
			WebURL            string      `json:"web_url"`
			AvatarURL         interface{} `json:"avatar_url"`
			GitSSHURL         string      `json:"git_ssh_url"`
			GitHTTPURL        string      `json:"git_http_url"`
			Namespace         string      `json:"namespace"`
			VisibilityLevel   int         `json:"visibility_level"`
			PathWithNamespace string      `json:"path_with_namespace"`
			DefaultBranch     string      `json:"default_branch"`
			CiConfigPath      string      `json:"ci_config_path"`
			Homepage          string      `json:"homepage"`
			URL               string      `json:"url"`
			SSHURL            string      `json:"ssh_url"`
			HTTPURL           string      `json:"http_url"`
		} `json:"target"`
		LastCommit struct {
			ID        string    `json:"id"`
			Message   string    `json:"message"`
			Timestamp time.Time `json:"timestamp"`
			URL       string    `json:"url"`
			Author    struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
		} `json:"last_commit"`
		WorkInProgress      bool        `json:"work_in_progress"`
		TotalTimeSpent      int         `json:"total_time_spent"`
		HumanTotalTimeSpent interface{} `json:"human_total_time_spent"`
		HumanTimeEstimate   interface{} `json:"human_time_estimate"`
	} `json:"object_attributes"`
	Labels     interface{} `json:"labels"`
	Repository struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
	} `json:"repository"`
}

type SystemRepositoryUpdatePayload struct {
	EventName  string `json:"event_name"`
	UserID     int    `json:"user_id"`
	UserName   string `json:"user_name"`
	UserEmail  string `json:"user_email"`
	UserAvatar string `json:"user_avatar"`
	ProjectID  int    `json:"project_id"`
	Project    struct {
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebURL            string      `json:"web_url"`
		AvatarURL         interface{} `json:"avatar_url"`
		GitSSHURL         string      `json:"git_ssh_url"`
		GitHTTPURL        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		Homepage          string      `json:"homepage"`
		URL               string      `json:"url"`
		SSHURL            string      `json:"ssh_url"`
		HTTPURL           string      `json:"http_url"`
	} `json:"project"`
	Changes []struct {
		Before string `json:"before"`
		After  string `json:"after"`
		Ref    string `json:"ref"`
	} `json:"changes"`
	Refs []string `json:"refs"`
}