export class Message{
    constructor(ID, Source, Destination, Content){
        this.ID = ID
        this.Source = Source
        this.Destination = Destination
        this.Content = Content
        this.Time = new Date()
    }
}
