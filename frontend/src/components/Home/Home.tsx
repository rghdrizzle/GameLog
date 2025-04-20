import RecommendedCarousal from "./components/RecommendedCarousal";
import TopGamesCarousal from "./components/TopGamesCarousal";

export default function Home(){
    return (
        <div className="flex justify-center items-center flex-col gap-10">
            {/* home page's top/new games */}
            <TopGamesCarousal />
            {/* home page's recommended games */}
            <RecommendedCarousal />
        </div>
    )
}