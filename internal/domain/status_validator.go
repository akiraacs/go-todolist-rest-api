package domain

func (status TaskStatus) IsValidStatus() bool {
    switch status {
    case StatusPending, StatusInProgress, StatusDone:
        return true
    default:
        return false
    }
}