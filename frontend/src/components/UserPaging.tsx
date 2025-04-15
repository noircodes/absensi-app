import { useLoaderData } from 'react-router';

export default function UserPaging() {

  let users = useLoaderData()
  console.log(users)
  return (
    <>
      <h1>USER PAGING</h1>
    </>
  )
}