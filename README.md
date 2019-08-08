# psutilsql

CLI tool that can be processed by SQL using  [gopsutil](https://github.com/shirou/gopsutil) library.

SQL input/output is handled by [trdsql](https://github.com/noborus/trdsql).
Therefore, CSV, JSON, LTSV, MarkDown, Raw, Vertical, and TBLN can be selected as the output format.

## install

```console
$ go get -u github.com/shirou/gopsutil...
```

## Usage

```console
$ psutilsql command
```

### SQL

The query command(\<query\> can be omitted) can execute SQL.

```console
$ psutilsql query "SELECT Total,Used,Free FROM virtualmemory"
or     
$ psutilsql "SELECT Total,Used,Free FROM virtualmemory"       

+-------------+------------+------------+
|    Total    |    Used    |    Free    |
+-------------+------------+------------+
| 16687091712 | 6468083712 | 2399399936 |
+-------------+------------+------------+
```

#### Table list

List of table names that can be used.

Displayed with the following command:
```console
$ psutilsql table
```

|      name       |
|-----------------|
| cpuinfo         |
| cpupercent      |
| cputime         |
| diskpartition   |
| diskusage       |
| docker          |
| hostinfo        |
| hosttemperature |
| hostuser        |
| loadavg         |
| loadmisc        |
| net             |
| process         |
| processex       |
| swapmemory      |
| virtualmemory   |


### Command

Display values using command and options without using SQL.

```console
$ psutilsql host --users
+---------+----------+------+------------+
|  User   | Terminal | Host |  Started   |
+---------+----------+------+------------+
| noborus | tty7     | :0   | 1564096509 |
+---------+----------+------+------------+
```

```console
$ psutilsql --help
SQL for running processes and system utilization.

SQL can be executed on the information acquired using gopsutil library.
Default SQL is provided, so you can omit SQL if you select a command.

Usage:
  psutilsql [flags]
  psutilsql [command]

Available Commands:
  completion  Generates bash/zsh completion scripts
  cpu         CPU information
  disk        DISK information
  docker      docker information
  help        Help about any command
  host        host information
  load        load information
  mem         memory information
  net         net information
  process     process information
  query       SQL query command
  table       table list

Flags:
  -d, --Delimiter string   output header (CSV only)
  -O, --Header             output header (CSV only)
  -o, --OutFormat string   output format
  -q, --Query string       query
  -h, --help               help for psutilsql
  -t, --toggle             Help message for toggle
```