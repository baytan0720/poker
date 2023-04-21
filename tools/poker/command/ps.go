/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package command

import (
	"context"
	"errors"
	"fmt"
	"poker/pkg/alert"
	"poker/pkg/service"
	"poker/tools/poker/util"
	"time"

	"github.com/spf13/cobra"
)

// PsCmd represents the ps command
var PsCmd = &cobra.Command{
	Use:    "ps",
	Short:  "List containers",
	Args:   cobra.NoArgs,
	Run:    ps,
	PreRun: util.Connect,
}

func init() {
	PsCmd.Flags().BoolP("all", "a", false, "Show all containers (default shows just running)")
	PsCmd.Flags().BoolP("quiet", "q", false, "Only display container IDs")
	PsCmd.Flags().BoolP("detail", "d", false, "More Detail of container")
}

func ps(cmd *cobra.Command, _ []string) {
	r, err := util.Client.PsContainer(context.Background(), &service.PsContainersReq{})
	if err != nil {
		alert.Error(err)
	}

	if r.Status != 0 {
		alert.Error(errors.New(r.Msg))
	}

	all, _ := cmd.Flags().GetBool("all")
	quiet, _ := cmd.Flags().GetBool("quiet")
	detail, _ := cmd.Flags().GetBool("detail")

	if !quiet {
		if detail {
			printDetailTitle()
		} else {
			printTitle()
		}
	}
	for _, container := range r.Containers {
		if container.State.Status == "Exited" && !all {
			continue
		}
		if quiet {
			alert.Println(container.Id[:12])
			continue
		}
		if detail {
			printPsDetail(container)
		} else {
			printPs(container)
		}
	}
}

func printTitle() {
	fmt.Printf("CONTAINER ID\tIMAGE\tCOMMAND\t\tCREATED\t\tSTATUS\t\t\tNAME\n")
}

func printDetailTitle() {
	fmt.Printf("CONTAINER ID\tIMAGE\tCOMMAND\t\tCREATED\t\tSTATUS\tNAME\tPID\tSTART\t\tFINISH\t\tERROR\n")
}

func printPs(container *service.ContainerInfo) {
	fmt.Printf("%s\t%s\t%s\t%s\t%s %s\t%s\n",
		container.Id[:12],
		container.Image,
		cutCommand(container.Command),
		since(container.Created),
		container.State.Status,
		StartOrFinish(container.State.Status, container.Created, container.State.Start, container.State.Finish),
		container.Name)
}

func printPsDetail(container *service.ContainerInfo) {
	fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%d\t%s\t%s\t%s\n",
		container.Id[:12],
		container.Image,
		cutCommand(container.Command),
		since(container.Created),
		container.State.Status,
		container.Name,
		container.State.Pid,
		since(container.State.Start),
		since(container.State.Finish),
		container.State.Error)

}

func since(t int64) string {
	if t == 0 {
		return ""
	}
	since := time.Duration(time.Now().UnixNano() - t)
	if since < time.Minute {
		return fmt.Sprintf("%d second ago", int(since.Seconds()))
	}
	if since < time.Hour {
		return fmt.Sprintf("%d minites ago", int(since.Minutes()))
	}
	if since < 24*time.Hour {
		return fmt.Sprintf("%d hours ago", int(since.Hours()))
	}
	if since < 7*24*time.Hour {
		return fmt.Sprintf("%d days ago", int(since.Hours()/24))
	}
	if since < 30*24*time.Hour {
		return fmt.Sprintf("%d weeks ago", int(since.Hours()/24/7))
	}
	if since < 12*30*24*time.Hour {
		return fmt.Sprintf("%d months ago", int(since.Hours()/24/30))
	}
	return fmt.Sprintf("%d years ago", int(since.Hours()/24/30/12))
}

func StartOrFinish(status string, created, start, finish int64) string {
	if status == "Running" {
		return since(start)
	} else if status == "Created" {
		return since(created)
	} else {
		return since(finish)
	}
}

func cutCommand(cmd string) string {
	if len(cmd) < 16 {
		return cmd + "\t"
	} else {
		return cmd[:13] + "..."
	}
}
