package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("To print in colors in a Print we have to write \\033 and",
		"then the following codes:\n",
		" * \033[0m[0m (normal/reset)\033[0m\n",
		" * \033[1m[1m\033[0m\n",
		" * \033[2m[2m\033[0m\n",
		" * \033[3m[3m\033[0m\n",
		" * \033[4m[4m\033[0m & \033[21m[21m\033[0m\n",
		" * \033[5m[5m\033[0m & \033[6m[6m\033[0m\n",
		" * \033[7m[7m\033[0m\n",
		" * \033[8m[8m\033[0m (invisible, can be copied)\n",
		" * \033[9m[9m\033[0m\n",
		" * \033[30m[30m\033[0m \033[30m\033[41m[30m & [41m\033[0m (very dark letters, change background to see)\n",
		" * \033[31m[31m\033[0m\n",
		" * \033[32m[32m\033[0m\n",
		" * \033[33m[33m\033[0m\n",
		" * \033[34m[34m\033[0m\n",
		" * \033[35m[35m\033[0m\n",
		" * \033[36m[36m\033[0m\n",
		" * \033[37m[37m\033[0m\n",
		" * \033[41m[41m\033[0m\n",
		" * \033[42m[42m\033[0m\n",
		" * \033[43m[43m\033[0m\n",
		" * \033[44m[44m\033[0m\n",
		" * \033[45m[45m\033[0m\n",
		" * \033[46m[46m\033[0m\n",
		" * \033[47m[47m\033[0m\n",
		" *      \033[53m[53m\033[0m (moved to the right to see the line)\n",
		" * \033[90m[90m\033[0m\n",
		" * \033[91m[91m\033[0m\n",
		" * \033[92m[92m\033[0m\n",
		" * \033[93m[93m\033[0m\n",
		" * \033[94m[94m\033[0m\n",
		" * \033[95m[95m\033[0m\n",
		" * \033[96m[96m\033[0m\n",
		" * \033[100m[100m\033[0m\n",
		" * \033[101m[101m\033[0m\n",
		" * \033[102m[102m\033[0m\n",
		" * \033[103m[103m\033[0m\n",
		" * \033[104m[104m\033[0m\n",
		" * \033[105m[105m\033[0m\n",
		" * \033[106m[106m\033[0m\n",
		" * \033[107m[107m\033[0m\n",
	)

	fmt.Println("And now, a loop:")
	for i := 0; i <= 110; i++ {
		fmt.Print("*** \033[" + strconv.Itoa(i) + "m[" + strconv.Itoa(i) + "\033[0m ")
	}
	fmt.Println()
}
