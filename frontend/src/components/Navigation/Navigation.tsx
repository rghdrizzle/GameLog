export default function Sidebar() {
  return (
    <div className="flex justify-between items-center gap-20">
      {/* sidebar's header */}
      <div className="flex justify-center items-center gap-2">
        <img src="/public/icons/game_white.svg" />
        <h2 className="text-3xl">GameLog.</h2>
      </div>

      {/* sidebar's content */}
      <div className="flex justify-center items-start gap-10">
        <div className="flex justify-center items-center gap-2">
          <p>Home</p>
        </div>

        <div className="flex justify-center items-center gap-2">
          <p>Categories</p>
        </div>

        <div className="flex justify-center items-center gap-2">
          <p>Community</p>
        </div>
      </div>
    </div>
  );
}
