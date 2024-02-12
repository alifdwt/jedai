import { Button } from "@/components/ui/button";
import {
  ArrowDownToLine,
  ArrowRightFromLine,
  CheckCircle,
  Leaf,
} from "lucide-react";
import Link from "next/link";
import { Oswald } from "next/font/google";
import { cn } from "@/lib/utils";

const oswald = Oswald({
  weight: "700",
  subsets: ["latin"],
});

const perks = [
  {
    name: "Instant Delivery",
    Icon: ArrowDownToLine,
    description:
      "Get your assets delivered to your email in seconds and download them right away.",
  },
  {
    name: "Guaranteed Quality",
    Icon: CheckCircle,
    description:
      "Every asset on our platform is verified by our team to ensure our highest quality standards. Not happy? We offer a 30-day refund guarantee.",
  },
  {
    name: "For the Planet",
    Icon: Leaf,
    description:
      "We've pledged 1% of sales to the preservation and restoration of the natural environment.",
  },
];

export default function Home() {
  return (
    <>
      <div className="py-20 bg-daintree-950">
        <div className="text-center flex flex-col items-center mx-auto max-w-3xl">
          <h3 className="px-5 py-3 bg-daintree-900 rounded-full text-white uppercase">
            lorem ipsum dolor sit amet
          </h3>
          <h1
            className={cn(
              "mt-6 text-4xl font-bold sm:text-6xl uppercase text-white",
              oswald.className
            )}
          >
            Your marketplace for high-quality{" "}
            <span className="text-secondary">digital assets</span>.
          </h1>
          <p className="mt-6 text-lg max-w-prose text-muted-foreground">
            Welcome to DigitalHippo. Every asset on our platform is verified by
            our team to ensure our highest quality standards.
          </p>
          <div className="flex flex-col sm:flex-row gap-4 mt-6">
            <Button asChild variant={"secondary"} className="rounded-full px-5">
              <Link href="/products" className="flex gap-2">
                <p>Telusuri</p> <ArrowRightFromLine size={18} />
              </Link>
            </Button>
            <Button
              asChild
              className="bg-daintree-200 text-primary rounded-full px-5 hover:bg-daintree-300"
            >
              <Link href="/products" className="flex gap-2">
                <p>Coba gratis</p> <ArrowRightFromLine size={18} />
              </Link>
            </Button>
          </div>
        </div>
      </div>

      <section className="border-t border-gray-200 bg-gray-50">
        <div className="mx-auto w-full max-w-screen-xl px-2.5 md:px-20 py-20">
          <div className="grid grid-cols-1 gap-y-12 sm:grid-cols-2 sm:gap-x-6 lg:grid-cols-3 lg:gap-x-8 lg:gap-y-0">
            {perks.map((perk) => (
              <div
                key={perk.name}
                className="text-center md:flex md:items-start md:text-left lg:block lg:text-center"
              >
                <div className="md:flex-shrink-0 flex justify-center">
                  <div className="h-16 w-16 flex items-center justify-center rounded-full bg-daintree-200 text-daintree-800">
                    {<perk.Icon className="w-1/3 h-1/3" />}
                  </div>
                </div>

                <div className="mt-6 md:ml-4 md:mt-0 lg:ml-0 lg:mt-6">
                  <h3 className="text-base font-medium text-gray-900">
                    {perk.name}
                  </h3>
                  <p className="mt-3 text-sm text-muted-foreground">
                    {perk.description}
                  </p>
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>
    </>
  );
}
