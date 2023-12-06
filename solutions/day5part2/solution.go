package day5part2

import (
	"aoc2023/solutions/day5part1"
	"aoc2023/utils"
	"log"
	"strconv"
)

type Solver struct{}

type Range struct {
	start int
	end   int
}

func (r Range) has(value int) bool {
	return value >= r.start && value <= r.end
}

func GetMappedValue(destination int, maps []day5part1.Map) int {
	for _, m := range maps {
		if destination >= m.Destination && destination < m.Destination+m.Length {
			return m.Source + destination - m.Destination
		}
	}
	return destination
}

func seedInRanges(seedRanges []Range, seed int) bool {
	for _, r := range seedRanges {
		if r.has(seed) {
			return true
		}
	}
	return false
}

func getLowestLocation(almanac day5part1.Almanac) int {
	var location int
	var seedRanges []Range
	for i := 0; i < len(almanac.Seeds); i += 2 {
		seedRanges = append(seedRanges, Range{start: almanac.Seeds[i], end: almanac.Seeds[i] + almanac.Seeds[i+1] - 1})
	}
	log.Printf("SeedRanges: %-v", seedRanges)
	for i := 0; true; i++ {
		hum := GetMappedValue(i, almanac.HumidityToLocation)
		temp := GetMappedValue(hum, almanac.TemperatureToHumidity)
		light := GetMappedValue(temp, almanac.LightToTemperature)
		water := GetMappedValue(light, almanac.WaterToLight)
		fert := GetMappedValue(water, almanac.FertilizerToWater)
		soil := GetMappedValue(fert, almanac.SoilToFertilizer)
		seed := GetMappedValue(soil, almanac.SeedToSoil)
		// log.Printf("Seed: %d, Soil: %d, Fert: %d, Water: %d, Light: %d, Temp: %d, Hum: %d, Loc: %d", seed, soil, fert, water, light, temp, hum, i)
		if seedInRanges(seedRanges, seed) {
			location = i
			break
		}

	}
	return location
}

func (Solver) Solve(input string) string {
	lines := utils.GetLines(input)
	almanac := day5part1.ParseAlmanac(lines)
	return strconv.Itoa(getLowestLocation(almanac))
}

//
//    Everyone will starve if you only plant such a small number of seeds.
//    Re-reading the almanac, it looks like the seeds: line actually
//    describes ranges of seed numbers.
//
//    The values on the initial seeds: line come in pairs. Within each pair,
//    the first value is the start of the range and the second value is the
//    length of the range. So, in the first line of the example above:
// seeds: 79 14 55 13
//
//    This line describes two ranges of seed numbers to be planted in the
//    garden. The first range starts with seed number 79 and contains 14
//    values: 79, 80, ..., 91, 92. The second range starts with seed number
//    55 and contains 13 values: 55, 56, ..., 66, 67.
//
//    Now, rather than considering four seed numbers, you need to consider a
//    total of 27 seed numbers.
//
//    In the above example, the lowest location number can be obtained from
//    seed number 82, which corresponds to soil 84, fertilizer 84, water 84,
//    light 77, temperature 45, humidity 46, and location 46. So, the lowest
//    location number is 46.
//
//    Consider all of the initial seed numbers listed in the ranges on the
//    first line of the almanac. What is the lowest location number that
//    corresponds to any of the initial seed numbers?
//
//    Answer: ____________________ [Submit]
//
//    Although it hasn't changed, you can still get your puzzle input.
//
//    You can also [Shareon Twitter Mastodon] this puzzle.
