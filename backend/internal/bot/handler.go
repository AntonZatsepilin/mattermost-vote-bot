package bot

func (b *MattermostBot) handleMessage(postID string) {
    post, resp := b.Client.GetPost(postID, "")
    if resp.Error != nil {
        b.Logger.Error(resp.Error)
        return
    }
    
    if post.UserId == b.User.Id {
        return
    }
    
    switch {
    case strings.HasPrefix(post.Message, "/poll create"):
        b.handleCreatePoll(post)
    case strings.HasPrefix(post.Message, "/poll vote"):
        b.handleVote(post)
    case strings.HasPrefix(post.Message, "/poll results"):
        b.handleResults(post)
    }
}

func (b *MattermostBot) handleCreatePoll(post *model.Post) {
    poll, err := b.Service.CreatePoll(...)
}