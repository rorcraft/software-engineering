 /*
	* public class LogEntry {
	*  public final long startTime; // start time of a job in millisec granularity
	*  public final long endTime; // end time of a job in millisec granularity.
	*  public final long ram; // the amount of ram the job occupies.
	*  public final long jobId;
	*  ... constructor ...
	* }
	*
	* running total of RAM
	* |
	* |             3GB
	* |           -----
	* |      2GB
	* |     ------
	* | 1GB                       -----------
	* |-----           -----------
	* |
	* |____________________________________________________time
	*
	* Find the peakRAM when the input is a collection of LogEntry objects
	*/

package main
import (
	"fmt"
	"sort"
)

type LogEntry struct {
	startTime int
	endTime int
	ram int
	jobId int
}

type RamLog struct {
	time int
	ram int
}

type RamLogs []RamLog

func (r RamLogs) Len() int { return len(r) }
func (r RamLogs) Less(i, j int) bool {
	return r[i].time < r[j].time || (r[i].time == r[j].time && r[i].ram < r[j].ram)
}
func (r RamLogs) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

func findPeakRam(logEntries []LogEntry) int {
	peakRam, curRam := 0, 0
	// break log entries into startTime +ram + endTime -ram
	ramLogs := make([]RamLog, 2 * len(logEntries))
	for i, entry := range logEntries {
		ii := 2 * i
		ramLogs[ii] = RamLog{time: entry.startTime, ram: entry.ram }
		ramLogs[ii+1] = RamLog{time: entry.endTime, ram: -entry.ram }
	}
	// radix sort maybe faster
	sort.Sort(RamLogs(ramLogs))
	// loop log and keep score of curRam, and keep peak.
	for _, ramlog := range ramLogs {
		// if starttime = endTime, makes the sort comparison more complicated 
		// if time is the same, remove ram for end time first.
		fmt.Println(ramlog.time)
		curRam = curRam + ramlog.ram
		if (curRam > peakRam) {
			peakRam = curRam
		}
	}

	return peakRam
}

func main() {
	logs := make([]LogEntry, 5)
	logs[0] = LogEntry{startTime: 10, endTime: 12, ram: 1 }
	logs[1] = LogEntry{startTime: 10, endTime: 13, ram: 2 }
	logs[2] = LogEntry{startTime: 12, endTime: 13, ram: 2 }
	logs[3] = LogEntry{startTime: 13, endTime: 15, ram: 2 }
	peakRam := findPeakRam(logs)
	fmt.Println("peakRam", peakRam)
}
