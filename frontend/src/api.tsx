export const fetchUserById = async (id: string) => {
    const res = await fetch(`http://localhost:8080/api/v1/user/${id}`)
    console.log(`http://localhost:8080/api/v1/user/${id}`)
    if (!res.ok) {
        throw new Error('Failed to fetch user')
    }
    return res.json()
}