# Task Manager CLI

A practical command-line task management application built with GuruUI. This application demonstrates all the enhanced features of the GuruUI library including colored status messages, icons, npm-style components, and interactive workflows.

## Features

### 🎯 **Task Management**
- **Add Tasks**: Create new tasks with title, description, priority, and due dates
- **List Tasks**: View all tasks in a clean, bordered table format
- **View Details**: See comprehensive task information with priority color coding
- **Update Status**: Change task status (Pending → In Progress → Completed)
- **Delete Tasks**: Remove tasks with confirmation prompts
- **Dashboard**: View task statistics and progress overview

### 🎨 **Enhanced UI Components**
- **Colored Status Messages**: 
  - ℹ️ Blue for info
  - ✓ Green for success  
  - ⚠️ Yellow for warnings
  - ✗ Red for errors
- **NPM-style Components**: Package-like formatting for tasks
- **Interactive Menus**: Colored prompts and numbered selections
- **Progress Bars**: Visual completion tracking
- **Tables**: Clean, bordered data display
- **Priority Colors**: Visual priority indicators

### 📊 **Smart Features**
- **Due Date Tracking**: Automatic overdue detection
- **Priority Levels**: Low, Medium, High, Critical with color coding
- **Status Workflow**: Logical task progression
- **Statistics Dashboard**: Real-time task metrics
- **Progress Tracking**: Overall completion rates

## Screenshots

The application features a clean, professional interface with:

- **Main Menu**: Organized action selection
- **Task Tables**: Structured data display with borders
- **Progress Bars**: Visual completion indicators
- **Status Messages**: Color-coded feedback
- **Interactive Forms**: User-friendly input prompts

## Usage

### Building and Running

```bash
# Build the application
go build -o task-manager main.go

# Run the application
./task-manager
```

### Using the Makefile

```bash
# Build all examples including task manager
make examples

# Run just the task manager
make tasks

# Run all examples
make all-examples
```

### Application Flow

1. **Start**: Welcome message and main menu
2. **Add Task**: 
   - Enter title and description
   - Select priority level
   - Set optional due date
3. **Manage Tasks**: List, view, update, or delete existing tasks
4. **Monitor Progress**: Use dashboard to track completion rates
5. **Exit**: Session summary and goodbye

## Code Structure

### Core Components

- **Task Struct**: Represents individual task items
- **TaskManager**: Handles all task operations
- **UI Integration**: Uses GuruUI for all interface elements

### Key Methods

- `AddTask()`: Creates new tasks with validation
- `ListTasks()`: Displays tasks in table format
- `ViewTask()`: Shows detailed task information
- `UpdateTaskStatus()`: Changes task workflow status
- `DeleteTask()`: Removes tasks with confirmation
- `ShowDashboard()`: Displays statistics and progress

### UI Features Demonstrated

- **Status Messages**: Info, Success, Warn, Error with icons
- **Interactive Selection**: Numbered option menus
- **Data Tables**: Bordered data display
- **Progress Bars**: Visual progress indicators
- **NPM Components**: Package-style formatting
- **Color Coding**: Priority and status indicators

## Technical Details

### Dependencies
- **GuruUI**: Terminal UI library (local dependency)
- **Go Standard Library**: Time, strings, strconv, fmt

### Architecture
- **Clean Separation**: Business logic separate from UI
- **Error Handling**: Graceful error messages and validation
- **Data Persistence**: In-memory storage (can be extended to file/database)
- **User Experience**: Intuitive workflows and clear feedback

### Extensibility
The application is designed to be easily extended with:
- File-based storage
- Database integration
- Task categories and tags
- Reminder notifications
- Export/import functionality
- Team collaboration features

## Learning Value

This application demonstrates:

1. **Real-world GuruUI Usage**: Practical implementation of all library features
2. **CLI Application Design**: Professional command-line interface patterns
3. **User Experience**: Intuitive workflows and clear feedback
4. **Code Organization**: Clean separation of concerns
5. **Error Handling**: Graceful validation and user feedback
6. **Interactive Design**: Engaging user experience in terminal

## Future Enhancements

Potential improvements include:
- **Data Persistence**: Save tasks to file or database
- **Task Categories**: Organize tasks by project or type
- **Due Date Reminders**: Notify users of upcoming deadlines
- **Task Templates**: Predefined task structures
- **Export Options**: Generate reports in various formats
- **Keyboard Shortcuts**: Faster navigation and actions

## Conclusion

The Task Manager CLI is a comprehensive demonstration of GuruUI's capabilities in a real-world application. It shows how the library can be used to create professional, user-friendly command-line interfaces that rival modern GUI applications in terms of usability and visual appeal.

Perfect for developers learning GuruUI or anyone looking for a practical task management solution!
