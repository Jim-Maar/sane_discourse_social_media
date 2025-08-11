package types

type ReactionType string

const (
	ReactionTypeAgree             ReactionType = "agree"
	ReactionTypeStrongAgree       ReactionType = "strong_agree"
	ReactionTypeDisagree          ReactionType = "disagree"
	ReactionTypeStrongDisagree    ReactionType = "strong_disagree"
	ReactionTypeImportant         ReactionType = "important"
	ReactionTypeStrongImportant   ReactionType = "strong_important"
	ReactionTypeUnimportant       ReactionType = "unimportant"
	ReactionTypeStrongUnimportant ReactionType = "strong_unimportant"
	ReactionTypeUpvote            ReactionType = "upvote"
	ReactionTypeStrongUpvote      ReactionType = "strong_upvote"
	ReactionTypeDownvote          ReactionType = "downvote"
	ReactionTypeStrongDownvote    ReactionType = "strong_downvote"
)
