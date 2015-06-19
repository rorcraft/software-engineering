## Hadoop

* NameNode - keeps the directory tree of all files in the file system, and tracks where across the cluster the file data is kept. It does not store the data of these files itself.
* DataNode - stores data, responds to requests from the NameNode for filesystem operations. (DataNode and TaskTracker should be together for locality)
* JobTracker - Client submits job to JobTracker. Talks to NameNode, locates TaskTracker nodes, submits work to TaskTracker.
* TaskTracker - configured with a set of slots, these indicate the number of tasks that it can accept. Accepts tasks - Map, Reduce and Shuffle operations - from a JobTracker.
