package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"vyshkat/domain"
	"vyshkat/fleet"
	"vyshkat/reports"
	"vyshkat/servicecenter"

	"go.uber.org/dig"
)

func printCountByKind(report reports.Reports, park *fleet.Fleet) {
	items := park.All()
	counts := report.CountByKind(items)

	kinds := make([]string, 0, len(counts))
	for kind := range counts {
		kinds = append(kinds, kind)
	}
	sort.Strings(kinds)

	fmt.Println("Всего на балансе:", len(items))
	for _, kind := range kinds {
		fmt.Printf("  %s: %d\n", kind, counts[kind])
	}
}

func printBeginnerFriendly(report reports.Reports, park *fleet.Fleet) {
	list := report.BeginnerFriendly(park.Vehicles())
	fmt.Println("Всего подходят новичкам:", len(list))
	for _, item := range list {
		fmt.Printf("%s - %s - %s\n", item.Name(), item.Kind(), item.InventoryNumber())
	}
}

func printInventory(report reports.Reports, park *fleet.Fleet) {
	list := report.SortedByInventoryNumber(park.All())
	fmt.Println("Всего в парке:", len(list))
	for _, item := range list {
		fmt.Printf("%s - %s - %s\n", item.Name(), item.Kind(), item.InventoryNumber())
	}
}

func accept(park *fleet.Fleet, v domain.Rentable) {
	if err := park.AddVehicle(v); err != nil {
		fmt.Printf("ОТКЛОНЁН  %s: %v\n", v.Name(), err)
		return
	}
	fmt.Printf("принят    %s (%s)\n", v.Name(), v.InventoryNumber())
}

func main() {
	container := dig.New()

	must(container.Provide(servicecenter.New, dig.As(new(fleet.Inspector))))
	must(container.Provide(func() domain.BeginnerPolicy {
		return domain.NewBeginnerPolicy(6)
	}))
	must(container.Provide(fleet.NewFleet))
	must(container.Provide(reports.New))

	must(container.Invoke(run))
}

func must(err error) {
	if err != nil {
		fmt.Println("ошибка сборки зависимостей:", err)
		os.Exit(1)
	}
}

func run(park *fleet.Fleet, report reports.Reports) {
	fmt.Println("=== Приёмка техники ===")
	accept(park, domain.NewScooter("Самокат-1", "VK-0001", 7))
	accept(park, domain.NewElectricScooter("Электросамокат-1", "VK-0002", 8, 1.2))
	accept(park, domain.NewBicycle("Велосипед-1", "VK-0003", 7))
	accept(park, domain.NewElectricBicycle("Электровелосипед-1", "VK-0004", 6, 2.1))
	accept(park, domain.NewScooter("Самокат-сложный", "VK-0005", 4))

	if err := park.AddVehicle(domain.NewScooter("Самокат-брак", "VK-0009", 15)); err != nil {
		fmt.Println("ОТКЛОНЁН техосмотром:", err)
	}

	if err := park.AddVehicle(domain.NewBicycle("Велосипед-брак", "", 7)); err != nil {
		fmt.Println("ОТКЛОНЁН техосмотром:", err)
	}

	if err := park.AddVehicle(domain.NewElectricBicycle("Электровелосипед-брак", "VK-0010", 6, -3.1)); err != nil {
		fmt.Println("ОТКЛОНЁН техосмотром:", err)
	}

	park.AddEquipment(domain.NewHelmet("Шлем-1", "VK-0100"))
	park.AddEquipment(domain.NewChargingCabinet("Зарядный шкаф-1", "VK-0200", 4.5))

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println()
		fmt.Println("=== Парк «ВышКат» ===")
		fmt.Println("1. Число единиц по видам")
		fmt.Println("2. Суммарное потребление энергии")
		fmt.Println("3. Техника для новичков")
		fmt.Println("4. Инвентаризация")
		fmt.Println("0. Выход")
		fmt.Print("> ")

		if !scanner.Scan() {
			return // ввод закончился
		}
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			printCountByKind(report, park)
		case "2":
			fmt.Printf("Суммарное потребление: %.2f кВт·ч\n", report.TotalEnergyKWh(park.All()))
		case "3":
			printBeginnerFriendly(report, park)
		case "4":
			printInventory(report, park)
		case "0":
			return
		default:
			fmt.Println("Нет такого пункта, введите цифру из списка")
		}
	}
}
