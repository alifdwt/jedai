const UserPage = ({ params }: { params: { userId: string } }) => {
  return (
    <div>
      <h3>Selamat datang di halaman user</h3>
      <h1 className="text-2xl">{params.userId}</h1>
    </div>
  );
};

export default UserPage;
