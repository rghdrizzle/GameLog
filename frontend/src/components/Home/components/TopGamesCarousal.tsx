import { Card, CardContent } from "@/components/ui/card";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "@/components/ui/carousel";

export default function TopGamesCarousal({}) {
  return (
    <div className="flex flex-col w-full gap-5">
      <div className="flex justify-center items-start flex-col px-4 ml-8">
        <h2
          className="text-2xl"
          style={{ fontFamily: "sohne_bold" }}
        >
          Top Games
        </h2>
        <p className="text-sm">Most popular games right now</p>
      </div>

      <div className="flex items-center justify-center w-full">
        <Carousel className="w-[90vw]">
          <CarouselContent className="h-full">
            {Array.from({ length: 1}).map((_, index) => (
              <CarouselItem key={index} className="h-[30vh] w-full">
                  <Card className="h-full py-0">
                    <CardContent className="h-full px-0">
                      <div className="flex items-center justify-center h-full w-full">
                        {index + 1}
                      </div>
                    </CardContent>
                  </Card>
              </CarouselItem>
            ))}
          </CarouselContent>
          <CarouselPrevious />
          <CarouselNext />
        </Carousel>
      </div>
    </div>
  );
}
