import { useLoaderData } from 'react-router';

export default function UserDetail() {

  let user = useLoaderData() as { id: string; name: string; }
  console.log(user)
  return (
    <div>
      <h2>{user.name}</h2>
      <p>ID: {user.id}</p>
    </div>
  );
}