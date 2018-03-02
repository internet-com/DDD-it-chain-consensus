package consensus

type ConsensusID struct{
	ID string
}

func NewConsensusID (id string) ConsensusID{
	return ConsensusID{
		ID: id,
	}
}

type LeaderID string

type Consensus struct {
	ConsensusID  ConsensusID
	Parliament    Parliament
	Block        Block
	CurrentState State
}

func NewConsensus(id ConsensusID,parliament Parliament,block Block) *Consensus{

	return &Consensus{
		ConsensusID:  id,
		Parliament:   parliament,
		Block:        block,
		CurrentState: new(IdleState)}
}

func (c *Consensus) Start(){
	c.CurrentState = new(PreprepareState)
}