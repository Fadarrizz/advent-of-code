use std::fs;
use uuid::Uuid;

#[derive(Debug, PartialEq, Clone)]
struct Node {
    id: Uuid,
    name: String,
    size: usize,
    children: Vec<Node>,
    parent: Uuid,
    is_dir: bool,
}

impl Node {
    fn new_dir(parent: Uuid, name: &str) -> Node {
        Node {
            id: Uuid::new_v4(),
            name: String::from(name),
            size: 0,
            children: Vec::new(),
            parent,
            is_dir: true
        }
    }

    fn new_file(parent: Uuid, name: &str, size: usize) -> Node {
        Node {
            id: Uuid::new_v4(),
            name: String::from(name),
            size,
            children: Vec::new(),
            parent,
            is_dir: false
        }
    }

    fn add(&mut self, child: Node) {
        self.children.push(child);
    }

    fn find_by_name(&mut self, name: &str) -> Option<&mut Node> {
        if self.name == name {
            return Some(self);
        }

        for child in &mut self.children {
            if child.name == name {
                return Some(child);
            }
        }

        None
    }

    fn find(&mut self, id: Uuid) -> Option<&mut Node> {
        if self.id == id {
            return Some(self);
        }

        for child in &mut self.children {
            let found = child.find(id);

            if found.is_some() {
                return found;
            }
        }

        None
    }

    fn get_size(&mut self) -> usize {
        if self.size > 0 {
            return self.size;
        }

        let mut acc = 0;
        for child in &mut self.children {
            acc += child.get_size();
        }

        self.size = acc;

        acc
    }
}

fn parse(s: &str) -> Node {
    let mut seen_dirs: Vec<Uuid> = Vec::new();
    let mut root_node = Node::new_dir(Uuid::nil(), "/");
    let mut current_node = &mut root_node;

    for line in s.lines() {
        if line == "$ cd /" || line == "$ ls" {
            continue;
        }

        else if line.starts_with("dir") {
            let dir = Node::new_dir(current_node.id.clone(), &line[4..]);
            current_node.add(dir);
        }

        else if line.starts_with("$ cd") {
            let dir_name = &line[5..];

            if dir_name == ".." {
                let prev_node = root_node.find(seen_dirs.pop().unwrap()).unwrap();
                current_node = prev_node;
            } else {
                seen_dirs.push(current_node.id.clone());
                current_node = current_node.find_by_name(dir_name).unwrap();
            }
        }

        else {
            let (size, name) = line.split_once(" ").unwrap();
            let file = Node::new_file(
                current_node.id.clone(), 
                name, 
                size.parse::<usize>().unwrap()
            );

            current_node.add(file);
        }
    }

    root_node.get_size();

    root_node
}

fn part1(mut root_node: Node) -> usize {
    let mut acc = 0;
    let mut nodes: Vec<Uuid> = vec!(root_node.id.clone());

    while nodes.len() > 0 {
        let node = root_node.find(nodes.pop().unwrap()).unwrap();

        if node.is_dir && node.size <= 100_000 {
            acc += node.size;
        }

        for child in &node.children {
            nodes.push(child.id.clone());
        }
    }

    acc
}

fn part2(mut root_node: Node) -> usize {
    let unused = 70_000_000 - root_node.size;
    let needed = 30_000_000 - unused;

    let mut nodes: Vec<Uuid> = vec!(root_node.id.clone());
    let mut sizes: Vec<usize> = Vec::new();
    while nodes.len() > 0 {
        let node = root_node.find(nodes.pop().unwrap()).unwrap();
    
        if node.size > needed {
            sizes.push(node.size);
        }

        for child in &node.children {
            nodes.push(child.id.clone());
        }
    }

    *sizes.iter().min().unwrap()
}

fn main() {
    let input = fs::read_to_string("./input.txt").unwrap();

    let node = parse(&input);

    // println!("{}", part1(node));
    println!("{}", part2(node));
}

#[test]
fn test_find_child() {
    let mut node = Node::new_dir(Uuid::nil(), "/");

    let child1 = Node::new_file("/", "file1", 0);
    let child2 = Node::new_file("/", "file2", 0);

    node.add(child1);
    node.add(child2);

    assert_eq!("file1", node.find(child1.id.clone()).unwrap().name);
    assert_eq!("file2", node.find(child2.id.clone()).unwrap().name);
}

#[test]
fn test_get_size() {
    let mut node = Node::new_dir("/", "/");

    let child1 = Node::new_file("/", "file1", 10);
    let child2 = Node::new_file("/", "file1", 10);

    node.add(child1);
    node.add(child2);

    assert_eq!(20, node.get_size());
}
