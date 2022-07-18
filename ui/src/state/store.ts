export interface SharedState<T> {
    data: T
    loading: boolean
    error: string
    fetch(): Promise<any>
}

export interface State<T> {
    data: T
    loading: boolean
    error: string
}
