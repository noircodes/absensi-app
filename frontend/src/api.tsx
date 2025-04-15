export const fetchUserById = async (id: string) => {
    const res = await fetch(`http://localhost:8080/api/v1/user/get/${id}`)
    console.log(`http://localhost:8080/api/v1/user/get/${id}`)
    if (!res.ok) {
        throw new Error('Failed to fetch user')
    }
    return res.json()
}

export const fetchUsers = async () => {
    const res = await fetch(`http://localhost:8080/api/v1/user/all`)
    console.log(`http://localhost:8080/api/v1/user/all`)
    if (!res.ok) {
        throw new Error('Failed to fetch users')
    }
    return res.json()
}